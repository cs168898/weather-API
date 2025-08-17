# Weather API

A weather API written in Go, featuring Redis caching, third-party weather data integration, and a RESTful API structure.

## Overview

This project provides a complete weather API service that demonstrates how to:

- Initialize and use Redis cache for storing weather data
- Integrate with third-party weather APIs (Visual Crossing Weather)
- Implement RESTful endpoints with Gin framework
- Handle caching strategies for improved performance
- Structure a Go project with proper separation of concerns

## Features

- **Redis Caching**: Intelligent caching of weather data with TTL (Time To Live)
- **Weather Data API**: Integration with Visual Crossing Weather API
- **RESTful Endpoints**: Clean API structure using Gin framework
- **Location-based Queries**: Get weather data for specific locations
- **Performance Optimization**: Cache-first approach to reduce API calls
- **Error Handling**: Graceful error handling for invalid locations or API issues

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) installed (version 1.19+)
- [Redis](https://redis.io/download) server running locally (default: `localhost:6379`)
- Visual Crossing Weather API key (sign up at [Visual Crossing](https://www.visualcrossing.com/))

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/weather-API.git
   cd weather-API
   ```

2. Install Go dependencies:
   ```sh
   go mod tidy
   ```

3. Create a `.env` file in the root directory:
   ```env
   WEATHER_API=your_visual_crossing_api_key_here
   ```

### Running the API

1. Start your Redis server if not already running:
   ```sh
   redis-server
   ```

2. Run the API:
   ```sh
   go run main.go
   ```

3. The API will start on `localhost:8080`

## API Endpoints

### GET `/getWeather/:location`

Retrieves weather data for a specific location.

**Parameters:**
- `location` (path parameter): The city/location name (e.g., "london", "new-york")

**Response:**
- Returns weather data in JSON format
- If cached data exists, returns from cache
- If no cache, fetches from Visual Crossing Weather API and caches the result

**Example:**
```sh
curl http://localhost:8080/getWeather/london
```

## Project Structure

```
Golang-Weather-API/
├── main.go                 # Application entry point
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksums
├── .env                    # Environment variables (create this)
├── models/
│   └── cacheObjects.go    # Data structures for caching
├── redis/
│   ├── redis.go           # Redis connection and cache initialization
│   ├── checkRedisCache.go # Cache retrieval functions
│   └── setCache.go        # Cache storage functions
├── router/
│   ├── runRouter.go       # Router setup and endpoint configuration
│   └── handlers/
│       └── getWeatherFunction.go # Weather API endpoint handler
└── utils/
    └── thirdPartyResponse.go # Third-party API integration utilities
```

### Key Components

- **`main.go`**: Initializes Redis cache, loads environment variables, and starts the router
- **`redis/`**: Handles all Redis operations including cache initialization, storage, and retrieval
- **`router/`**: Manages API routing and endpoint configuration using Gin framework
- **`handlers/`**: Contains the business logic for processing API requests
- **`models/`**: Defines data structures used throughout the application
- **`utils/`**: Contains utility functions for external API calls

## Caching Strategy

The API implements an intelligent caching strategy:

1. **Cache Key Format**: `{location},{date1},{date2}` (currently hardcoded to 2025-08-15)
2. **TTL**: 1 hour for cached weather data
3. **Cache-First Approach**: Checks Redis before making external API calls
4. **Automatic Population**: New data is automatically cached after external API calls

## Configuration

- **Redis**: Configured to run on `localhost:6379`
- **API Port**: Server runs on port `8080`
- **Weather API**: Uses Visual Crossing Weather API with metric units
- **Cache TTL**: 1 hour for weather data

## Error Handling

- Graceful handling of invalid locations
- API key validation
- Cache miss fallback to external API
- JSON response formatting for all scenarios

## Dependencies

Key Go packages used:
- `github.com/gin-gonic/gin` - Web framework
- `github.com/go-redis/redis/v8` - Redis client
- `github.com/go-redis/cache/v8` - Redis caching layer
- `github.com/joho/godotenv` - Environment variable loading

## Future Enhancements

- [ ] Add date range support for historical weather data
- [ ] Implement multiple weather data providers
- [ ] Add rate limiting and API key management
- [ ] Implement weather data filtering and formatting
- [ ] Add metrics and monitoring
- [ ] Support for multiple cache strategies
- [ ] Add unit tests and integration tests



Roadmap.sh backend projects:
https://roadmap.sh/projects/weather-api-wrapper-service