// Code generated by generator, DO NOT EDIT.
package trttlcache

import (
	trcache "github.com/RangelReale/trcache"
	"time"
)

func WithDefaultDuration[K comparable, V any](duration time.Duration) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptDefaultDuration(duration)
			return true
		}
		return false
	})
}
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
func WithValidator[K comparable, V any](validator trcache.Validator[V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptValidator(validator)
			return true
		}
		return false
	})
}
func WithGetTouch[K comparable, V any](touch bool) trcache.GetOption {
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case GetOptions[K, V]:
			opt.OptTouch(touch)
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
func (ob *RootOptionBuilder[K, V]) WithDefaultDuration(duration time.Duration) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithDefaultDuration[K, V](duration))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithName(name string) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithName[K, V](name))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithValidator(validator trcache.Validator[V]) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithValidator[K, V](validator))
	return ob
}

type GetOptionBuilder[K comparable, V any] struct {
	trcache.GetOptionBuilderBase
}

func GetOpt[K comparable, V any]() *GetOptionBuilder[K, V] {
	return &GetOptionBuilder[K, V]{}
}
func (ob *GetOptionBuilder[K, V]) WithGetTouch(touch bool) *GetOptionBuilder[K, V] {
	ob.AppendOptions(WithGetTouch[K, V](touch))
	return ob
}

type rootOptions[K comparable, V any] struct {
	trcache.IsRootOptionsImpl
	callDefaultDeleteOptions []trcache.DeleteOption
	callDefaultGetOptions    []trcache.GetOption
	callDefaultSetOptions    []trcache.SetOption
	defaultDuration          time.Duration
	name                     string
	validator                trcache.Validator[V]
}

var _ Options[string, string] = &rootOptions[string, string]{}

func (o *rootOptions[K, V]) OptCallDefaultDeleteOptions(options ...trcache.DeleteOption) {
	o.callDefaultDeleteOptions = options
}
func (o *rootOptions[K, V]) OptCallDefaultGetOptions(options ...trcache.GetOption) {
	o.callDefaultGetOptions = options
}
func (o *rootOptions[K, V]) OptCallDefaultSetOptions(options ...trcache.SetOption) {
	o.callDefaultSetOptions = options
}
func (o *rootOptions[K, V]) OptDefaultDuration(duration time.Duration) {
	o.defaultDuration = duration
}
func (o *rootOptions[K, V]) OptName(name string) {
	o.name = name
}
func (o *rootOptions[K, V]) OptValidator(validator trcache.Validator[V]) {
	o.validator = validator
}

type getOptions[K comparable, V any] struct {
	trcache.IsGetOptionsImpl
	customOptions []interface{}
	touch         bool
}

var _ GetOptions[string, string] = &getOptions[string, string]{}

func (o *getOptions[K, V]) OptCustomOptions(customOptions []interface{}) {
	o.customOptions = customOptions
}
func (o *getOptions[K, V]) OptTouch(touch bool) {
	o.touch = touch
}

type setOptions[K comparable, V any] struct {
	trcache.IsSetOptionsImpl
	duration time.Duration
}

var _ SetOptions[string, string] = &setOptions[string, string]{}

func (o *setOptions[K, V]) OptDuration(duration time.Duration) {
	o.duration = duration
}

type deleteOptions[K comparable, V any] struct {
	trcache.IsDeleteOptionsImpl
}

var _ DeleteOptions[string, string] = &deleteOptions[string, string]{}
