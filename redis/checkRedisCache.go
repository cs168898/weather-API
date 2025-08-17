package redis

import (
	"context"
	models "main/models"
)

func CheckRedisCache(key string) models.Object {

	ctx := context.TODO()

	var wanted models.Object
	if err := myCache.Get(ctx, key, &wanted); err == nil {
		// return the value if found in the cache
		return wanted
	}

	// return an empty object if not found in the cache
	return models.Object{}

}
