package main

import (
	"github.com/gin-gonic/gin"
)

func runRouter() {
	// create a new gin router with the getWeather endpoint
	router := gin.Default()
	router.GET("/getWeather/:location", getWeatherFunction)

	router.Run("localhost:8080")

}
