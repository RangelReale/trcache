package trredis

import (
	"time"

	"github.com/RangelReale/trcache"
)

// Option

type CacheOptions[K comparable, V any] interface {
	trcache.IsOption
	trcache.CacheFnDefaultOptions[K, V]
	OptName(string)
	OptKeyCodec(trcache.KeyCodec[K])
	OptValueCodec(trcache.Codec[V])
	OptValidator(trcache.Validator[V])
	OptDefaultDuration(time.Duration)
	OptGetFunc(GetFunc[K, V])
	OptSetFunc(SetFunc[K, V])
	OptDelFunc(DelFunc[K, V])
}

type cacheOptions[K comparable, V any] struct {
	trcache.IsOptionImpl
	fnDefaultGet    []trcache.GetOption[K, V]
	fnDefaultSet    []trcache.SetOption[K, V]
	fnDefaultDelete []trcache.DeleteOption[K, V]
	name            string
	keyCodec        trcache.KeyCodec[K]
	valueCodec      trcache.Codec[V]
	validator       trcache.Validator[V]
	defaultDuration time.Duration
	getFunc         GetFunc[K, V]
	setFunc         SetFunc[K, V]
	delFunc         DelFunc[K, V]
}

var _ CacheOptions[string, string] = &cacheOptions[string, string]{}

func (c *cacheOptions[K, V]) OptFnDefaultGetOpt(i []trcache.GetOption[K, V]) {
	c.fnDefaultGet = i
}

func (c *cacheOptions[K, V]) OptFnDefaultSetOpt(i []trcache.SetOption[K, V]) {
	c.fnDefaultSet = i
}

func (c *cacheOptions[K, V]) OptFnDefaultDeleteOpt(i []trcache.DeleteOption[K, V]) {
	c.fnDefaultDelete = i
}

func (c *cacheOptions[K, V]) OptName(s string) {
	c.name = s
}

func (c *cacheOptions[K, V]) OptKeyCodec(t trcache.KeyCodec[K]) {
	c.keyCodec = t
}

func (c *cacheOptions[K, V]) OptValueCodec(t trcache.Codec[V]) {
	c.valueCodec = t
}

func (c *cacheOptions[K, V]) OptValidator(t trcache.Validator[V]) {
	c.validator = t
}

func (c *cacheOptions[K, V]) OptDefaultDuration(duration time.Duration) {
	c.defaultDuration = duration
}

func (c *cacheOptions[K, V]) OptGetFunc(fn GetFunc[K, V]) {
	c.getFunc = fn
}

func (c *cacheOptions[K, V]) OptSetFunc(fn SetFunc[K, V]) {
	c.setFunc = fn
}

func (c *cacheOptions[K, V]) OptDelFunc(fn DelFunc[K, V]) {
	c.delFunc = fn
}

func WithName[K comparable, V any](name string) trcache.Option[K, V] {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptName(name)
			return true
		}
		return false
	})
}

func WithKeyCodec[K comparable, V any](keyCodec trcache.KeyCodec[K]) trcache.Option[K, V] {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptKeyCodec(keyCodec)
			return true
		}
		return false
	})
}

func WithValueCodec[K comparable, V any](valueCodec trcache.Codec[V]) trcache.Option[K, V] {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptValueCodec(valueCodec)
			return true
		}
		return false
	})
}

func WithValidator[K comparable, V any](validator trcache.Validator[V]) trcache.Option[K, V] {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptValidator(validator)
			return true
		}
		return false
	})
}

func WithDefaultDuration[K comparable, V any](defaultDuration time.Duration) trcache.Option[K, V] {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptDefaultDuration(defaultDuration)
			return true
		}
		return false
	})
}

// Cache get options

type CacheGetOptions[K comparable, V any] interface {
	trcache.IsGetOption
	trcache.GetOptions[K, V]
	OptCustomParams(any)
	OptGetFunc(GetFunc[K, V])
}

type cacheGetOptions[K comparable, V any] struct {
	trcache.IsGetOptionImpl
	customOptions []any
	customParams  any
	getFunc       GetFunc[K, V]
}

var _ CacheGetOptions[string, string] = &cacheGetOptions[string, string]{}

func (c *cacheGetOptions[K, V]) OptCustomOptions(anies []any) {
	c.customOptions = anies
}

func (c *cacheGetOptions[K, V]) OptCustomParams(params any) {
	c.customParams = params
}

func (c *cacheGetOptions[K, V]) OptGetFunc(fn GetFunc[K, V]) {
	c.getFunc = fn
}

func WithCacheGetCustomParam[K comparable, V any](param any) trcache.GetOption[K, V] {
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheGetOptions[K, V]:
			opt.OptCustomParams(param)
			return true
		}
		return false
	})
}

func WithCacheGetGetFunc[K comparable, V any](fn GetFunc[K, V]) trcache.GetOption[K, V] {
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheGetOptions[K, V]:
			opt.OptGetFunc(fn)
			return true
		}
		return false
	})
}

// Cache set options

type CacheSetOptions[K comparable, V any] interface {
	trcache.IsSetOption
	trcache.SetOptions[K, V]
	OptCustomParams(any)
	OptSetFunc(SetFunc[K, V])
}

type cacheSetOptions[K comparable, V any] struct {
	trcache.IsSetOptionImpl
	duration     time.Duration
	customParams any
	setFunc      SetFunc[K, V]
}

var _ CacheSetOptions[string, string] = &cacheSetOptions[string, string]{}

func (c *cacheSetOptions[K, V]) OptDuration(duration time.Duration) {
	c.duration = duration
}

func (c *cacheSetOptions[K, V]) OptCustomParams(customParams any) {
	c.customParams = customParams
}

func (c *cacheSetOptions[K, V]) OptSetFunc(fn SetFunc[K, V]) {
	c.setFunc = fn
}

func WithCacheSetSetFunc[K comparable, V any](fn SetFunc[K, V]) trcache.SetOption[K, V] {
	return trcache.SetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheSetOptions[K, V]:
			opt.OptSetFunc(fn)
			return true
		}
		return false
	})
}

// Cache delete options

type CacheDeleteOptions[K comparable, V any] interface {
	trcache.IsDeleteOption
	trcache.DeleteOptions[K, V]
	OptCustomParams(any)
	OptDelFunc(DelFunc[K, V])
}

type cacheDeleteOptions[K comparable, V any] struct {
	trcache.IsDeleteOptionImpl
	customParams any
	delFunc      DelFunc[K, V]
}

var _ CacheDeleteOptions[string, string] = &cacheDeleteOptions[string, string]{}

func (c *cacheDeleteOptions[K, V]) OptCustomParams(customParams any) {
	c.customParams = customParams
}

func (c *cacheDeleteOptions[K, V]) OptDelFunc(fn DelFunc[K, V]) {
	c.delFunc = fn
}

func WithCacheDeleteDelFunc[K comparable, V any](fn DelFunc[K, V]) trcache.DeleteOption[K, V] {
	return trcache.DeleteOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheDeleteOptions[K, V]:
			opt.OptDelFunc(fn)
			return true
		}
		return false
	})
}
