package redis

import (
	"context"
	"fmt"
	"time"

	models "main/models"

	"github.com/go-redis/cache/v8"
)

func SetCache(key string, value models.Object) {
	fmt.Printf("Setting cache for key: %s\n", key)
	ctx := context.TODO()

	if err := myCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   time.Hour,
	}); err != nil {
		panic(err)
	}
}
