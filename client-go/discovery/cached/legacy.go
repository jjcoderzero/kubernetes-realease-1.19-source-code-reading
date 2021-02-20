package memory

import (
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
)

// NewMemCacheClient is DEPRECATED. Use memory.NewMemCacheClient directly.
func NewMemCacheClient(delegate discovery.DiscoveryInterface) discovery.CachedDiscoveryInterface {
	return memory.NewMemCacheClient(delegate)
}

// ErrCacheNotFound is DEPRECATED. Use memory.ErrCacheNotFound directly.
var ErrCacheNotFound = memory.ErrCacheNotFound
