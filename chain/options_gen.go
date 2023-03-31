// Code generated by generator, DO NOT EDIT.
package chain

import trcache "github.com/RangelReale/trcache"

func WithName[K comparable, V any](name string) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptName(name)
			return true
		}
		return false
	})
}
func WithRefreshFunc[K comparable, V any](refreshFunc trcache.CacheRefreshFunc[K, V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptRefreshFunc(refreshFunc)
			return true
		}
		return false
	})
}
func WithSetPreviousOnGet[K comparable, V any](setPreviousOnGet bool) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptSetPreviousOnGet(setPreviousOnGet)
			return true
		}
		return false
	})
}
func WithGetGetStrategy[K comparable, V any](getStrategy GetStrategy[K, V]) trcache.GetOption {
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case GetOptions[K, V]:
			opt.OptGetStrategy(getStrategy)
			return true
		}
		return false
	})
}
func WithGetSetOptions[K comparable, V any](options ...trcache.SetOption) trcache.GetOption {
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case GetOptions[K, V]:
			opt.OptSetOptions(options...)
			return true
		}
		return false
	})
}
func WithSetSetStrategy[K comparable, V any](setStrategy SetStrategy[K, V]) trcache.SetOption {
	return trcache.SetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case SetOptions[K, V]:
			opt.OptSetStrategy(setStrategy)
			return true
		}
		return false
	})
}
func WithDeleteDeleteStrategy[K comparable, V any](deleteStrategy DeleteStrategy[K, V]) trcache.DeleteOption {
	return trcache.DeleteOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case DeleteOptions[K, V]:
			opt.OptDeleteStrategy(deleteStrategy)
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
func (ob *RootOptionBuilder[K, V]) WithName(name string) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithName[K, V](name))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithRefreshFunc(refreshFunc trcache.CacheRefreshFunc[K, V]) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithRefreshFunc[K, V](refreshFunc))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithSetPreviousOnGet(setPreviousOnGet bool) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithSetPreviousOnGet[K, V](setPreviousOnGet))
	return ob
}

type GetOptionBuilder[K comparable, V any] struct {
	trcache.GetOptionBuilderBase
}

func GetOpt[K comparable, V any]() *GetOptionBuilder[K, V] {
	return &GetOptionBuilder[K, V]{}
}
func (ob *GetOptionBuilder[K, V]) WithGetGetStrategy(getStrategy GetStrategy[K, V]) *GetOptionBuilder[K, V] {
	ob.AppendOptions(WithGetGetStrategy[K, V](getStrategy))
	return ob
}
func (ob *GetOptionBuilder[K, V]) WithGetSetOptions(options ...trcache.SetOption) *GetOptionBuilder[K, V] {
	ob.AppendOptions(WithGetSetOptions[K, V](options...))
	return ob
}

type SetOptionBuilder[K comparable, V any] struct {
	trcache.SetOptionBuilderBase
}

func SetOpt[K comparable, V any]() *SetOptionBuilder[K, V] {
	return &SetOptionBuilder[K, V]{}
}
func (ob *SetOptionBuilder[K, V]) WithSetSetStrategy(setStrategy SetStrategy[K, V]) *SetOptionBuilder[K, V] {
	ob.AppendOptions(WithSetSetStrategy[K, V](setStrategy))
	return ob
}

type DeleteOptionBuilder[K comparable, V any] struct {
	trcache.DeleteOptionBuilderBase
}

func DeleteOpt[K comparable, V any]() *DeleteOptionBuilder[K, V] {
	return &DeleteOptionBuilder[K, V]{}
}
func (ob *DeleteOptionBuilder[K, V]) WithDeleteDeleteStrategy(deleteStrategy DeleteStrategy[K, V]) *DeleteOptionBuilder[K, V] {
	ob.AppendOptions(WithDeleteDeleteStrategy[K, V](deleteStrategy))
	return ob
}
