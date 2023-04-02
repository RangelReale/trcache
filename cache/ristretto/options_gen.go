// Code generated by troptgen. DO NOT EDIT.

package trristretto

import (
	trcache "github.com/RangelReale/trcache"
	"time"
)

type RootOption = trcache.RootOption

func WithCallDefaultDeleteOptions[K comparable, V any](options ...trcache.DeleteOption) RootOption {
	return trcache.WithCallDefaultDeleteOptions[K, V](options...)
}
func WithCallDefaultGetOptions[K comparable, V any](options ...trcache.GetOption) RootOption {
	return trcache.WithCallDefaultGetOptions[K, V](options...)
}
func WithCallDefaultSetOptions[K comparable, V any](options ...trcache.SetOption) RootOption {
	return trcache.WithCallDefaultSetOptions[K, V](options...)
}
func WithDefaultDuration[K comparable, V any](duration time.Duration) RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptDefaultDuration(duration)
			return true
		}
		return false
	})
}
func WithEventualConsistency[K comparable, V any](eventualConsistency bool) RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptEventualConsistency(eventualConsistency)
			return true
		}
		return false
	})
}
func WithName[K comparable, V any](name string) RootOption {
	return trcache.WithName[K, V](name)
}
func WithValidator[K comparable, V any](validator trcache.Validator[V]) RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptValidator(validator)
			return true
		}
		return false
	})
}
func WithValueCodec[K comparable, V any](valueCodec trcache.Codec[V]) RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptValueCodec(valueCodec)
			return true
		}
		return false
	})
}

type GetOption = trcache.GetOption

func WithGetCustomOptions[K comparable, V any](customOptions []interface{}) GetOption {
	return trcache.WithGetCustomOptions[K, V](customOptions)
}

type SetOption = trcache.SetOption

func WithSetCost[K comparable, V any](cost int64) SetOption {
	return trcache.SetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case setOptions[K, V]:
			opt.OptCost(cost)
			return true
		}
		return false
	})
}
func WithSetDuration[K comparable, V any](duration time.Duration) SetOption {
	return trcache.WithSetDuration[K, V](duration)
}

type DeleteOption = trcache.DeleteOption
type rootOptionsImpl[K comparable, V any] struct {
	trcache.IsRootOptionsImpl
	callDefaultDeleteOptions []trcache.DeleteOption
	callDefaultGetOptions    []trcache.GetOption
	callDefaultSetOptions    []trcache.SetOption
	defaultDuration          time.Duration
	eventualConsistency      bool
	name                     string
	validator                trcache.Validator[V]
	valueCodec               trcache.Codec[V]
}

var _ options[string, string] = &rootOptionsImpl[string, string]{}

func (o *rootOptionsImpl[K, V]) OptCallDefaultDeleteOptions(options ...trcache.DeleteOption) {
	o.callDefaultDeleteOptions = options
}
func (o *rootOptionsImpl[K, V]) OptCallDefaultGetOptions(options ...trcache.GetOption) {
	o.callDefaultGetOptions = options
}
func (o *rootOptionsImpl[K, V]) OptCallDefaultSetOptions(options ...trcache.SetOption) {
	o.callDefaultSetOptions = options
}
func (o *rootOptionsImpl[K, V]) OptDefaultDuration(duration time.Duration) {
	o.defaultDuration = duration
}
func (o *rootOptionsImpl[K, V]) OptEventualConsistency(eventualConsistency bool) {
	o.eventualConsistency = eventualConsistency
}
func (o *rootOptionsImpl[K, V]) OptName(name string) {
	o.name = name
}
func (o *rootOptionsImpl[K, V]) OptValidator(validator trcache.Validator[V]) {
	o.validator = validator
}
func (o *rootOptionsImpl[K, V]) OptValueCodec(valueCodec trcache.Codec[V]) {
	o.valueCodec = valueCodec
}

type getOptionsImpl[K comparable, V any] struct {
	trcache.IsGetOptionsImpl
	customOptions []interface{}
}

var _ getOptions[string, string] = &getOptionsImpl[string, string]{}

func (o *getOptionsImpl[K, V]) OptCustomOptions(customOptions []interface{}) {
	o.customOptions = customOptions
}

type setOptionsImpl[K comparable, V any] struct {
	trcache.IsSetOptionsImpl
	cost     int64
	duration time.Duration
}

var _ setOptions[string, string] = &setOptionsImpl[string, string]{}

func (o *setOptionsImpl[K, V]) OptCost(cost int64) {
	o.cost = cost
}
func (o *setOptionsImpl[K, V]) OptDuration(duration time.Duration) {
	o.duration = duration
}

type deleteOptionsImpl[K comparable, V any] struct {
	trcache.IsDeleteOptionsImpl
}

var _ deleteOptions[string, string] = &deleteOptionsImpl[string, string]{}
