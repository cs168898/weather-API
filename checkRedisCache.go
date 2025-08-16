package main

import (
	"context"
)

func checkRedisCache(key string) Object {

	ctx := context.TODO()

	var wanted Object
	if err := mycache.Get(ctx, key, &wanted); err == nil {
    // return the value if found in the cache
		return wanted
	}

  // return an empty object if not found in the cache
  return Object{}
  
}
