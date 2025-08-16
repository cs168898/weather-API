# Weather API

A simple weather API written in Go, featuring Redis caching and environment variable support.

## Overview

This project provides a basic structure for a weather API service. It demonstrates how to:

- Initialize and use a Redis cache (via go-redis and go-redis/cache)
- Load environment variables from a `.env` file
- Set up a placeholder for routing and API endpoints

## Features

- Redis caching setup for storing and retrieving weather data
- Environment variable loading with `godotenv`
- Modular structure for future expansion (e.g., adding endpoints, filtering data)

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) installed
- [Redis](https://redis.io/download) server running locally (default: `localhost:6379`)

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

3. (Optional) Create a `.env` file for environment variables.

### Running the API

Start your Redis server if not already running.

Run the API:
```sh
go run main.go
```

You should see:
```
Weather API Started
```

## Project Structure

- `main.go` — Initializes Redis cache, loads environment variables, and starts the router.
- (Planned) Router and API endpoint logic to handle weather queries.

## Usage

API endpoints are to be implemented. The current codebase sets up the infrastructure for caching and environment management.

## Roadmap

- [ ] Implement API endpoints for weather data queries
- [ ] Integrate Redis `set` and `get` for caching weather responses
- [ ] Filter and return only relevant weather data to users

