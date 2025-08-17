package main

import (
	"fmt"
	redis "main/redis"
	router "main/router"

	"github.com/joho/godotenv"
)

func main() {

	fmt.Print("Weather API Started\n")

	godotenv.Load()

	// initialize the redis cache
	redis.InitRedis()
	// start the router and set up API endpoints
	router.RunRouter()

	fmt.Print("Weather api Ended")
}
