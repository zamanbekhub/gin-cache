package memory

import "github.com/jellydator/ttlcache/v2"

type Option func(c *ttlcache.Cache)

func WithSize(size int) Option {
	return func(c *ttlcache.Cache) {
		c.SetCacheSizeLimit(size)
	}
}
