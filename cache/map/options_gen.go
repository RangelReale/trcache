// Code generated by troptgen. DO NOT EDIT.

package trmap

import (
	trcache "github.com/RangelReale/trcache"
	"time"
)

func WithValidator[K comparable, V any](validator trcache.Validator[V]) trcache.RootOption {
	const optionName = "github.com/RangelReale/trcache/cache/map/options.Validator"
	const optionHash = uint64(0x8e4bac8c480a9c30)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptValidator(validator)
			return true
		}
		return false
	}, optionName, optionHash)
}

type rootOptionsImpl[K comparable, V any] struct {
	callDefaultDeleteOptions []trcache.DeleteOption
	callDefaultGetOptions    []trcache.GetOption
	callDefaultSetOptions    []trcache.SetOption
	name                     string
	validator                trcache.Validator[V]
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
func (o *rootOptionsImpl[K, V]) OptName(name string) {
	o.name = name
}
func (o *rootOptionsImpl[K, V]) OptValidator(validator trcache.Validator[V]) {
	o.validator = validator
}

type getOptionsImpl[K comparable, V any] struct{}

var _ getOptions[string, string] = &getOptionsImpl[string, string]{}

type setOptionsImpl[K comparable, V any] struct {
	duration time.Duration
}

var _ setOptions[string, string] = &setOptionsImpl[string, string]{}

func (o *setOptionsImpl[K, V]) OptDuration(duration time.Duration) {
	o.duration = duration
}

type deleteOptionsImpl[K comparable, V any] struct{}

var _ deleteOptions[string, string] = &deleteOptionsImpl[string, string]{}