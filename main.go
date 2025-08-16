package main

import (
	"fmt"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

//type weatherData struct {
//	Address     string `json:"address"`
//	Datetime    string `json:"datetime"`
//	Description string `json:"description"`
//}

var weatherDatas = make(map[string]any)

var mycache *cache.Cache

func main() {

	fmt.Print("Weather API Started\n")

	godotenv.Load()

	// initialize the redis cache
	mycache = initRedis()
	// start the router and set up API endpoints
	runRouter()

	fmt.Print("Weather api Ended")
}

// this function returns the cache object address from the ”*cache.Cache' type
func initRedis() *cache.Cache {
	redisAddr := "localhost:6379"

	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": redisAddr,
		},
	})

	mycache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return mycache

}
