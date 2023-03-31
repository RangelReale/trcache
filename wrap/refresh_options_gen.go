// Code generated by generator, DO NOT EDIT.
package wrap

import trcache "github.com/RangelReale/trcache"

func WithCallDefaultRefreshOptions[K comparable, V any](p0 ...trcache.RefreshOption) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case WrapRefreshOptions[K, V]:
			opt.OptCallDefaultRefreshOptions(p0...)
			return true
		}
		return false
	})
}
func WithDefaultRefreshFunc[K comparable, V any](p0 trcache.CacheRefreshFunc[K, V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case WrapRefreshOptions[K, V]:
			opt.OptDefaultRefreshFunc(p0)
			return true
		}
		return false
	})
}
func WithRefreshData[K comparable, V any](p0 interface{}) trcache.RefreshOption {
	return trcache.RefreshOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case WrapRefreshRefreshOptions[K, V]:
			opt.OptData(p0)
			return true
		}
		return false
	})
}
func WithRefreshRefreshFunc[K comparable, V any](p0 trcache.CacheRefreshFunc[K, V]) trcache.RefreshOption {
	return trcache.RefreshOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case WrapRefreshRefreshOptions[K, V]:
			opt.OptRefreshFunc(p0)
			return true
		}
		return false
	})
}
func WithRefreshSetOptions[K comparable, V any](p0 []trcache.SetOption) trcache.RefreshOption {
	return trcache.RefreshOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case WrapRefreshRefreshOptions[K, V]:
			opt.OptSetOptions(p0)
			return true
		}
		return false
	})
}

type RootOptionBuilder[K comparable, V any] struct {
	trcache.RootOptionBuilderBase
}

func RootOpt[K comparable, V any]() *RootOptionBuilder[K, V] {
	return &RootOptionBuilder[K, V]{}
}
func (ob *RootOptionBuilder[K, V]) WithCallDefaultRefreshOptions(p0 ...trcache.RefreshOption) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithCallDefaultRefreshOptions[K, V](p0...))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithDefaultRefreshFunc(p0 trcache.CacheRefreshFunc[K, V]) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithDefaultRefreshFunc[K, V](p0))
	return ob
}

type RefreshOptionBuilder[K comparable, V any] struct {
	trcache.RefreshOptionBuilderBase
}

func RefreshOpt[K comparable, V any]() *RefreshOptionBuilder[K, V] {
	return &RefreshOptionBuilder[K, V]{}
}
func (ob *RefreshOptionBuilder[K, V]) WithRefreshData(p0 interface{}) *RefreshOptionBuilder[K, V] {
	ob.AppendOptions(WithRefreshData[K, V](p0))
	return ob
}
func (ob *RefreshOptionBuilder[K, V]) WithRefreshRefreshFunc(p0 trcache.CacheRefreshFunc[K, V]) *RefreshOptionBuilder[K, V] {
	ob.AppendOptions(WithRefreshRefreshFunc[K, V](p0))
	return ob
}
func (ob *RefreshOptionBuilder[K, V]) WithRefreshSetOptions(p0 []trcache.SetOption) *RefreshOptionBuilder[K, V] {
	ob.AppendOptions(WithRefreshSetOptions[K, V](p0))
	return ob
}