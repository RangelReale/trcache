// Code generated by generator, DO NOT EDIT.
package wrap

import trcache "github.com/RangelReale/trcache"

func WithWrapDefaultRefreshFunc[K comparable, V any](refreshFunc trcache.CacheRefreshFunc[K, V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case WrapRefreshOptions[K, V]:
			opt.OptDefaultRefreshFunc(refreshFunc)
			return true
		}
		return false
	})
}

type WrapRootOptionBuilder[K comparable, V any] struct {
	trcache.RootOptionBuilderBase
}

func RootOpt[K comparable, V any]() *WrapRootOptionBuilder[K, V] {
	return &WrapRootOptionBuilder[K, V]{}
}
func (ob *WrapRootOptionBuilder[K, V]) WithWrapDefaultRefreshFunc(refreshFunc trcache.CacheRefreshFunc[K, V]) *WrapRootOptionBuilder[K, V] {
	ob.AppendOptions(WithWrapDefaultRefreshFunc[K, V](refreshFunc))
	return ob
}

type wrapRefreshOptions[K comparable, V any] struct {
	trcache.IsRootOptionsImpl
	callDefaultRefreshOptions []trcache.RefreshOption
	defaultRefreshFunc        trcache.CacheRefreshFunc[K, V]
}

var _ WrapRefreshOptions[string, string] = &wrapRefreshOptions[string, string]{}

func (o *wrapRefreshOptions[K, V]) OptCallDefaultRefreshOptions(options ...trcache.RefreshOption) {
	o.callDefaultRefreshOptions = options
}
func (o *wrapRefreshOptions[K, V]) OptDefaultRefreshFunc(refreshFunc trcache.CacheRefreshFunc[K, V]) {
	o.defaultRefreshFunc = refreshFunc
}

type wrapRefreshRefreshOptions[K comparable, V any] struct {
	trcache.IsRefreshOptionsImpl
	data        interface{}
	refreshFunc trcache.CacheRefreshFunc[K, V]
	setOptions  []trcache.SetOption
}

var _ WrapRefreshRefreshOptions[string, string] = &wrapRefreshRefreshOptions[string, string]{}

func (o *wrapRefreshRefreshOptions[K, V]) OptData(data interface{}) {
	o.data = data
}
func (o *wrapRefreshRefreshOptions[K, V]) OptRefreshFunc(refreshFunc trcache.CacheRefreshFunc[K, V]) {
	o.refreshFunc = refreshFunc
}
func (o *wrapRefreshRefreshOptions[K, V]) OptSetOptions(options ...trcache.SetOption) {
	o.setOptions = options
}
