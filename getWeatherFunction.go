package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// function to get the weather data
func getWeatherFunction(c *gin.Context) {

	weatherAPIKey := os.Getenv("WEATHER_API")

	location := c.Param("location")
	location = strings.ToLower(location)
	//date1 := c.Param("date1")
	//date2 := c.Param("date2")

	// define the key for the redis cache using the location and dates
	key := fmt.Sprintf("%s,%s,%s", location, "2025-08-15", "2025-08-15")
	// check if the data is in the redis cache
	cacheResult := Object{}
	cacheResult = checkRedisCache(key)

	// check if cachresult is empty, if its not empty then use it
	// else call the third party api

	if len(cacheResult.WeatherDatas) != 0 {
		fmt.Printf("Using cache for location: %s\n", location)
		// if the data is in the cache, use it
		weatherDatas = cacheResult.WeatherDatas
	} else {
		fmt.Printf("Calling third party API for location: %s\n", location)
		// this is the url that we will GET request to
		url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/%s/%s?key=%s&unitGroup=%s", location, "2025-08-15", "2025-08-15", weatherAPIKey, "metric")

		weatherDatas = getThirdPartyResponse(url)

		// Check if response is empty
		if len(weatherDatas) == 0 {
			weatherDatas = map[string]any{
				"error": "Please check the location or API Key provided",
			}
		} else {
			// if something exists inside the data
			// set the cache with the key and the weather data
			setCache(key, Object{WeatherDatas: weatherDatas})
		}

	}

	c.IndentedJSON(http.StatusOK, weatherDatas)
}
