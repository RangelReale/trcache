// Code generated by troptgen. DO NOT EDIT.

package chain

import (
	trcache "github.com/RangelReale/trcache"
	"time"
)

// WithDefaultStrategyCallback sets a callback function to receive strategy results.
func WithDefaultStrategyCallback[K comparable, V any](callback StrategyCallback) trcache.RootOption {
	const optionName = "github.com/RangelReale/trcache/cache/chain/options.DefaultStrategyCallback"
	const optionHash = uint64(0xba694c293ad3e7ba)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptDefaultStrategyCallback(callback)
			return true
		}
		return false
	}, optionName, optionHash)
}

// WithDeleteStrategy sets the [DeleteStrategy] to use for the chain operation. The default is
// [DeleteStrategyDeleteAll].
func WithDeleteStrategy[K comparable, V any](deleteStrategy DeleteStrategy[K, V]) trcache.RootOption {
	const optionName = "github.com/RangelReale/trcache/cache/chain/options.DeleteStrategy"
	const optionHash = uint64(0x9611f18cf7185dfb)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptDeleteStrategy(deleteStrategy)
			return true
		}
		return false
	}, optionName, optionHash)
}

// WithGetStrategy sets the [GetStrategy] to use for the chain operation. The default is
// [GetStrategyGetFirstSetPrevious].
func WithGetStrategy[K comparable, V any](getStrategy GetStrategy[K, V]) trcache.RootOption {
	const optionName = "github.com/RangelReale/trcache/cache/chain/options.GetStrategy"
	const optionHash = uint64(0x6eeb406b2a0672b8)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptGetStrategy(getStrategy)
			return true
		}
		return false
	}, optionName, optionHash)
}

// WithSetStrategy sets the [SetStrategy] to use for the chain operation. The default is
// [SetStrategySetAll].
func WithSetStrategy[K comparable, V any](setStrategy SetStrategy[K, V]) trcache.RootOption {
	const optionName = "github.com/RangelReale/trcache/cache/chain/options.SetStrategy"
	const optionHash = uint64(0x1dcc48665967d4c)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptSetStrategy(setStrategy)
			return true
		}
		return false
	}, optionName, optionHash)
}

// WithGetSetOptions adds options to the [Cache.Set] call done after one of the [Cache.Get] function calls succeeds.
func WithGetSetOptions[K comparable, V any](options ...trcache.SetOption) trcache.GetOption {
	const optionName = "github.com/RangelReale/trcache/cache/chain/getOptions.SetOptions"
	const optionHash = uint64(0x20cdc9d4030ddb85)
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case getOptions[K, V]:
			opt.OptSetOptions(options...)
			return true
		}
		return false
	}, optionName, optionHash)
}

// WithGetStrategyCallback sets a callback function to receive strategy results.
func WithGetStrategyCallback[K comparable, V any](callback StrategyCallback) trcache.GetOption {
	const optionName = "github.com/RangelReale/trcache/cache/chain/getOptions.StrategyCallback"
	const optionHash = uint64(0xf083d5cb17c16a0d)
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case getOptions[K, V]:
			opt.OptStrategyCallback(callback)
			return true
		}
		return false
	}, optionName, optionHash)
}

// WithSetStrategyCallback sets a callback function to receive strategy results.
func WithSetStrategyCallback[K comparable, V any](callback StrategyCallback) trcache.SetOption {
	const optionName = "github.com/RangelReale/trcache/cache/chain/setOptions.StrategyCallback"
	const optionHash = uint64(0xcc37a90344a23759)
	return trcache.SetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case setOptions[K, V]:
			opt.OptStrategyCallback(callback)
			return true
		}
		return false
	}, optionName, optionHash)
}

// WithDeleteStrategyCallback sets a callback function to receive strategy results.
func WithDeleteStrategyCallback[K comparable, V any](callback StrategyCallback) trcache.DeleteOption {
	const optionName = "github.com/RangelReale/trcache/cache/chain/deleteOptions.StrategyCallback"
	const optionHash = uint64(0xe0f7c60406a6643e)
	return trcache.DeleteOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case deleteOptions[K, V]:
			opt.OptStrategyCallback(callback)
			return true
		}
		return false
	}, optionName, optionHash)
}

type rootOptionsImpl[K comparable, V any] struct {
	callDefaultDeleteOptions []trcache.DeleteOption
	callDefaultGetOptions    []trcache.GetOption
	callDefaultSetOptions    []trcache.SetOption
	defaultStrategyCallback  StrategyCallback
	deleteStrategy           DeleteStrategy[K, V]
	getStrategy              GetStrategy[K, V]
	name                     string
	setStrategy              SetStrategy[K, V]
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
func (o *rootOptionsImpl[K, V]) OptDefaultStrategyCallback(callback StrategyCallback) {
	o.defaultStrategyCallback = callback
}
func (o *rootOptionsImpl[K, V]) OptDeleteStrategy(deleteStrategy DeleteStrategy[K, V]) {
	o.deleteStrategy = deleteStrategy
}
func (o *rootOptionsImpl[K, V]) OptGetStrategy(getStrategy GetStrategy[K, V]) {
	o.getStrategy = getStrategy
}
func (o *rootOptionsImpl[K, V]) OptName(name string) {
	o.name = name
}
func (o *rootOptionsImpl[K, V]) OptSetStrategy(setStrategy SetStrategy[K, V]) {
	o.setStrategy = setStrategy
}

type getOptionsImpl[K comparable, V any] struct {
	setOptions       []trcache.SetOption
	strategyCallback StrategyCallback
}

var _ getOptions[string, string] = &getOptionsImpl[string, string]{}

func (o *getOptionsImpl[K, V]) OptSetOptions(options ...trcache.SetOption) {
	o.setOptions = options
}
func (o *getOptionsImpl[K, V]) OptStrategyCallback(callback StrategyCallback) {
	o.strategyCallback = callback
}

type setOptionsImpl[K comparable, V any] struct {
	duration         time.Duration
	strategyCallback StrategyCallback
}

var _ setOptions[string, string] = &setOptionsImpl[string, string]{}

func (o *setOptionsImpl[K, V]) OptDuration(duration time.Duration) {
	o.duration = duration
}
func (o *setOptionsImpl[K, V]) OptStrategyCallback(callback StrategyCallback) {
	o.strategyCallback = callback
}

type deleteOptionsImpl[K comparable, V any] struct {
	strategyCallback StrategyCallback
}

var _ deleteOptions[string, string] = &deleteOptionsImpl[string, string]{}

func (o *deleteOptionsImpl[K, V]) OptStrategyCallback(callback StrategyCallback) {
	o.strategyCallback = callback
}
