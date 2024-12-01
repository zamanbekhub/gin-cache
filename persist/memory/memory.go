package memory

import (
	"errors"
	"github.com/chenyahui/gin-cache/persist"
	"reflect"
	"time"

	"github.com/jellydator/ttlcache/v2"
)

// MemoryStore local memory cache store
type MemoryStore struct {
	cache *ttlcache.Cache
}

// NewMemoryStore allocate a local memory store with default expiration
func NewMemoryStore(defaultExpiration time.Duration, opts ...Option) *MemoryStore {
	cacheStore := ttlcache.NewCache()
	_ = cacheStore.SetTTL(defaultExpiration)

	// disable SkipTTLExtensionOnHit default
	cacheStore.SkipTTLExtensionOnHit(true)
	for _, opt := range opts {
		opt(cacheStore)
	}

	return &MemoryStore{
		cache: cacheStore,
	}
}

// Set put key value pair to memory store, and expire after expireDuration
func (c *MemoryStore) Set(key string, value interface{}, expireDuration time.Duration) error {
	return c.cache.SetWithTTL(key, value, expireDuration)
}

// Delete remove key in memory store, do nothing if key doesn't exist
func (c *MemoryStore) Delete(key string) error {
	return c.cache.Remove(key)
}

// Get key in memory store, if key doesn't exist, return ErrCacheMiss
func (c *MemoryStore) Get(key string, value interface{}) error {
	val, err := c.cache.Get(key)
	if errors.Is(err, ttlcache.ErrNotFound) {
		return persist.ErrCacheMiss
	}

	v := reflect.ValueOf(value)
	v.Elem().Set(reflect.ValueOf(val))
	return nil
}
