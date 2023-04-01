// Code generated by generator, DO NOT EDIT.
package trredis

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
func WithKeyCodec[K comparable, V any](keyCodec trcache.KeyCodec[K]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptKeyCodec(keyCodec)
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
func WithRedisDelFunc[K comparable, V any](redisDelFunc RedisDelFunc[K, V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptRedisDelFunc(redisDelFunc)
			return true
		}
		return false
	})
}
func WithRedisGetFunc[K comparable, V any](redisGetFunc RedisGetFunc[K, V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptRedisGetFunc(redisGetFunc)
			return true
		}
		return false
	})
}
func WithRedisSetFunc[K comparable, V any](redisSetFunc RedisSetFunc[K, V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptRedisSetFunc(redisSetFunc)
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
func WithValueCodec[K comparable, V any](valueCodec trcache.Codec[V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptValueCodec(valueCodec)
			return true
		}
		return false
	})
}
func WithGetCustomParams[K comparable, V any](customParams interface{}) trcache.GetOption {
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case GetOptions[K, V]:
			opt.OptCustomParams(customParams)
			return true
		}
		return false
	})
}
func WithGetRedisGetFunc[K comparable, V any](redisGetFunc RedisGetFunc[K, V]) trcache.GetOption {
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case GetOptions[K, V]:
			opt.OptRedisGetFunc(redisGetFunc)
			return true
		}
		return false
	})
}
func WithSetCustomParams[K comparable, V any](customParams interface{}) trcache.SetOption {
	return trcache.SetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case SetOptions[K, V]:
			opt.OptCustomParams(customParams)
			return true
		}
		return false
	})
}
func WithSetRedisSetFunc[K comparable, V any](redisSetFunc RedisSetFunc[K, V]) trcache.SetOption {
	return trcache.SetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case SetOptions[K, V]:
			opt.OptRedisSetFunc(redisSetFunc)
			return true
		}
		return false
	})
}
func WithDeleteCustomParams[K comparable, V any](customParams interface{}) trcache.DeleteOption {
	return trcache.DeleteOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case DeleteOptions[K, V]:
			opt.OptCustomParams(customParams)
			return true
		}
		return false
	})
}
func WithDeleteRedisDelFunc[K comparable, V any](redisDelFunc RedisDelFunc[K, V]) trcache.DeleteOption {
	return trcache.DeleteOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case DeleteOptions[K, V]:
			opt.OptRedisDelFunc(redisDelFunc)
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
func (ob *RootOptionBuilder[K, V]) WithKeyCodec(keyCodec trcache.KeyCodec[K]) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithKeyCodec[K, V](keyCodec))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithName(name string) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithName[K, V](name))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithRedisDelFunc(redisDelFunc RedisDelFunc[K, V]) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithRedisDelFunc[K, V](redisDelFunc))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithRedisGetFunc(redisGetFunc RedisGetFunc[K, V]) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithRedisGetFunc[K, V](redisGetFunc))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithRedisSetFunc(redisSetFunc RedisSetFunc[K, V]) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithRedisSetFunc[K, V](redisSetFunc))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithValidator(validator trcache.Validator[V]) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithValidator[K, V](validator))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithValueCodec(valueCodec trcache.Codec[V]) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithValueCodec[K, V](valueCodec))
	return ob
}

type GetOptionBuilder[K comparable, V any] struct {
	trcache.GetOptionBuilderBase
}

func GetOpt[K comparable, V any]() *GetOptionBuilder[K, V] {
	return &GetOptionBuilder[K, V]{}
}
func (ob *GetOptionBuilder[K, V]) WithGetCustomParams(customParams interface{}) *GetOptionBuilder[K, V] {
	ob.AppendOptions(WithGetCustomParams[K, V](customParams))
	return ob
}
func (ob *GetOptionBuilder[K, V]) WithGetRedisGetFunc(redisGetFunc RedisGetFunc[K, V]) *GetOptionBuilder[K, V] {
	ob.AppendOptions(WithGetRedisGetFunc[K, V](redisGetFunc))
	return ob
}

type SetOptionBuilder[K comparable, V any] struct {
	trcache.SetOptionBuilderBase
}

func SetOpt[K comparable, V any]() *SetOptionBuilder[K, V] {
	return &SetOptionBuilder[K, V]{}
}
func (ob *SetOptionBuilder[K, V]) WithSetCustomParams(customParams interface{}) *SetOptionBuilder[K, V] {
	ob.AppendOptions(WithSetCustomParams[K, V](customParams))
	return ob
}
func (ob *SetOptionBuilder[K, V]) WithSetRedisSetFunc(redisSetFunc RedisSetFunc[K, V]) *SetOptionBuilder[K, V] {
	ob.AppendOptions(WithSetRedisSetFunc[K, V](redisSetFunc))
	return ob
}

type DeleteOptionBuilder[K comparable, V any] struct {
	trcache.DeleteOptionBuilderBase
}

func DeleteOpt[K comparable, V any]() *DeleteOptionBuilder[K, V] {
	return &DeleteOptionBuilder[K, V]{}
}
func (ob *DeleteOptionBuilder[K, V]) WithDeleteCustomParams(customParams interface{}) *DeleteOptionBuilder[K, V] {
	ob.AppendOptions(WithDeleteCustomParams[K, V](customParams))
	return ob
}
func (ob *DeleteOptionBuilder[K, V]) WithDeleteRedisDelFunc(redisDelFunc RedisDelFunc[K, V]) *DeleteOptionBuilder[K, V] {
	ob.AppendOptions(WithDeleteRedisDelFunc[K, V](redisDelFunc))
	return ob
}

type rootOptions[K comparable, V any] struct {
	trcache.IsRootOptionsImpl
	callDefaultDeleteOptions []trcache.DeleteOption
	callDefaultGetOptions    []trcache.GetOption
	callDefaultSetOptions    []trcache.SetOption
	defaultDuration          time.Duration
	keyCodec                 trcache.KeyCodec[K]
	name                     string
	redisDelFunc             RedisDelFunc[K, V]
	redisGetFunc             RedisGetFunc[K, V]
	redisSetFunc             RedisSetFunc[K, V]
	validator                trcache.Validator[V]
	valueCodec               trcache.Codec[V]
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
func (o *rootOptions[K, V]) OptKeyCodec(keyCodec trcache.KeyCodec[K]) {
	o.keyCodec = keyCodec
}
func (o *rootOptions[K, V]) OptName(name string) {
	o.name = name
}
func (o *rootOptions[K, V]) OptRedisDelFunc(redisDelFunc RedisDelFunc[K, V]) {
	o.redisDelFunc = redisDelFunc
}
func (o *rootOptions[K, V]) OptRedisGetFunc(redisGetFunc RedisGetFunc[K, V]) {
	o.redisGetFunc = redisGetFunc
}
func (o *rootOptions[K, V]) OptRedisSetFunc(redisSetFunc RedisSetFunc[K, V]) {
	o.redisSetFunc = redisSetFunc
}
func (o *rootOptions[K, V]) OptValidator(validator trcache.Validator[V]) {
	o.validator = validator
}
func (o *rootOptions[K, V]) OptValueCodec(valueCodec trcache.Codec[V]) {
	o.valueCodec = valueCodec
}

type getOptions[K comparable, V any] struct {
	trcache.IsGetOptionsImpl
	customOptions []interface{}
	customParams  interface{}
	redisGetFunc  RedisGetFunc[K, V]
}

var _ GetOptions[string, string] = &getOptions[string, string]{}

func (o *getOptions[K, V]) OptCustomOptions(customOptions []interface{}) {
	o.customOptions = customOptions
}
func (o *getOptions[K, V]) OptCustomParams(customParams interface{}) {
	o.customParams = customParams
}
func (o *getOptions[K, V]) OptRedisGetFunc(redisGetFunc RedisGetFunc[K, V]) {
	o.redisGetFunc = redisGetFunc
}

type setOptions[K comparable, V any] struct {
	trcache.IsSetOptionsImpl
	customParams interface{}
	duration     time.Duration
	redisSetFunc RedisSetFunc[K, V]
}

var _ SetOptions[string, string] = &setOptions[string, string]{}

func (o *setOptions[K, V]) OptCustomParams(customParams interface{}) {
	o.customParams = customParams
}
func (o *setOptions[K, V]) OptDuration(duration time.Duration) {
	o.duration = duration
}
func (o *setOptions[K, V]) OptRedisSetFunc(redisSetFunc RedisSetFunc[K, V]) {
	o.redisSetFunc = redisSetFunc
}

type deleteOptions[K comparable, V any] struct {
	trcache.IsDeleteOptionsImpl
	customParams interface{}
	redisDelFunc RedisDelFunc[K, V]
}

var _ DeleteOptions[string, string] = &deleteOptions[string, string]{}

func (o *deleteOptions[K, V]) OptCustomParams(customParams interface{}) {
	o.customParams = customParams
}
func (o *deleteOptions[K, V]) OptRedisDelFunc(redisDelFunc RedisDelFunc[K, V]) {
	o.redisDelFunc = redisDelFunc
}
