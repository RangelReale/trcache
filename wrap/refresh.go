package wrap

import (
	"context"
	"errors"

	"github.com/RangelReale/trcache"
)

type wrapRefreshCache[K comparable, V any, RD any] struct {
	options wrapRefreshOptionsImpl[K, V, RD]
	cache   trcache.Cache[K, V]
}

func NewWrapRefreshCache[K comparable, V any, RD any](cache trcache.Cache[K, V],
	options ...trcache.RootOption) (trcache.RefreshCache[K, V, RD], error) {
	ret := &wrapRefreshCache[K, V, RD]{
		cache: cache,
	}
	optErr := trcache.ParseRootOptions(&ret.options, options)
	if optErr != nil && !ret.options.ignoreOptionNotSupported {
		return nil, optErr
	}
	return ret, nil
}

func (c *wrapRefreshCache[K, V, RD]) Name() string {
	return c.cache.Name()
}

func (c *wrapRefreshCache[K, V, RD]) Get(ctx context.Context, key K,
	options ...trcache.GetOption) (V, error) {
	return c.cache.Get(ctx, key, options...)
}

func (c *wrapRefreshCache[K, V, RD]) Set(ctx context.Context, key K, value V,
	options ...trcache.SetOption) error {
	return c.cache.Set(ctx, key, value, options...)
}

func (c *wrapRefreshCache[K, V, RD]) Delete(ctx context.Context, key K,
	options ...trcache.DeleteOption) error {
	return c.cache.Delete(ctx, key, options...)
}

func (c *wrapRefreshCache[K, V, RD]) GetOrRefresh(ctx context.Context, key K, options ...trcache.RefreshOption) (V, error) {
	optns := wrapRefreshRefreshOptionsImpl[K, V, RD]{
		funcx: c.options.defaultRefreshFunc,
	}
	optErr := trcache.ParseRefreshOptions(&optns, c.options.callDefaultRefreshOptions, options)
	if optErr != nil && !optns.ignoreOptionNotSupported {
		var empty V
		return empty, optErr
	}

	ret, err := c.Get(ctx, key, optns.getOptions...)
	if err == nil {
		if c.options.metricsMetrics != nil {
			c.options.metricsMetrics.Hit(ctx, c.options.metricsName)
		}
		return ret, nil
	} else if err != nil && !errors.Is(err, trcache.ErrNotFound) {
		if c.options.metricsMetrics != nil {
			var cerr *trcache.CodecError
			if errors.As(err, &cerr) {
				c.options.metricsMetrics.Error(ctx, c.options.metricsName, trcache.MetricsErrorTypeDecode)
			} else {
				c.options.metricsMetrics.Error(ctx, c.options.metricsName, trcache.MetricsErrorTypeGet)
			}
		}
		var empty V
		return empty, err
	}

	if c.options.metricsMetrics != nil {
		c.options.metricsMetrics.Miss(ctx, c.options.metricsName)
	}

	// call refresh
	if optns.funcx == nil {
		var empty V
		return empty, errors.New("refresh function not set")
	}

	ret, err = optns.funcx(ctx, key, trcache.RefreshFuncOptions[RD]{
		Data: optns.data,
	})
	if err != nil {
		if c.options.metricsMetrics != nil {
			c.options.metricsMetrics.Error(ctx, c.options.metricsName, trcache.MetricsErrorTypeRefresh)
		}
		var empty V
		return empty, err
	}

	err = c.Set(ctx, key, ret, optns.setOptions...)
	if err != nil {
		if c.options.metricsMetrics != nil {
			var cerr *trcache.CodecError
			if errors.As(err, &cerr) {
				c.options.metricsMetrics.Error(ctx, c.options.metricsName, trcache.MetricsErrorTypeEncode)
			} else {
				c.options.metricsMetrics.Error(ctx, c.options.metricsName, trcache.MetricsErrorTypeSet)
			}
		}
		var empty V
		return empty, err
	}

	return ret, nil
}
