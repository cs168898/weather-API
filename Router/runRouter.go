package router

import (
	handlers "main/router/handlers"

	"github.com/gin-gonic/gin"
)

func RunRouter() {
	// create a new gin router with the getWeather endpoint
	router := gin.Default()
	router.GET("/getWeather/:location", handlers.GetWeatherFunction)

	router.Run("localhost:8080")

}
