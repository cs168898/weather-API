package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/cache/v8"
)

func setCache(key string, value Object) {
  fmt.Printf("Setting cache for key: %s\n", key)
	ctx := context.TODO()

	if err := mycache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   time.Hour,
	}); err != nil {
		panic(err)
	}
}
