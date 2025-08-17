package redis

import (
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var myCache *cache.Cache

// this function returns the cache object address from the ”*cache.Cache' type
func InitRedis() {
	redisAddr := "localhost:6379"

	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": redisAddr,
		},
	})

	// assign the new cache to MyCache variable
	myCache = cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

}
