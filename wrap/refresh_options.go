package wrap

import (
	"github.com/RangelReale/trcache"
)

// Option

// +troptgen root
type wrapRefreshOptions[K comparable, V any] interface {
	trcache.IsRootOptions
	trcache.MetricsOptions[K, V]
	trcache.CallDefaultRefreshOptions[K, V]
	OptDefaultRefreshFunc(refreshFunc trcache.CacheRefreshFunc[K, V])
}

// Cache refresh options

// +troptgen refresh
type wrapRefreshRefreshOptions[K comparable, V any] interface {
	trcache.IsRefreshOptions
	trcache.RefreshOptions[K, V]
}

//go:generate troptgen -prefix wrap
