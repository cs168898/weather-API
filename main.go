package main

import (
	"fmt"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"time"
	"os"
)

//type weatherData struct {
//	Address     string `json:"address"`
//	Datetime    string `json:"datetime"`
//	Description string `json:"description"`
//}

var weatherDatas = make(map[string]interface{})

var mycache *cache.Cache

func main() {
	fmt.Print("Weather API Started\n")

	// initialize the redis cache
	//initRedis()
	// start the router and set up API endpoints
	runRouter()

	fmt.Print("Weather api Ended")
}

// this function returns the cache object address from the ''*cache.Cache' type
func initRedis()(*cache.Cache) {
	redisURL := os.Getenv("REDIS_URL")
	
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			redisURL: ":6379",
			"server2": ":6380",
		},
	})

	mycache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return mycache

}
