# WeatherAPI CLI

A blazing-fast CLI tool written in Go to fetch real-time weather data using the [Visual Crossing API](https://www.visualcrossing.com/). Features Redis caching via Docker to avoid redundant API calls.

> ⚡ Built with Go, Redis, and Docker — made for devs who like their forecasts fast and terminal-native.

## Features

- 📡 Fetches current weather info for any location
- 💾 Caches results in Redis for quick repeated lookups  
- 🐳 Redis container via Docker Compose (no setup headaches)
- 🔧 Modular code with helper utilities

## Prerequisites

- [Go](https://golang.org/dl/) (v1.18+ recommended)
- [Docker](https://docs.docker.com/get-docker/)
- Visual Crossing API Key

## Setup

### 1. Clone the repository

```bash
git clone https://github.com/A-noob-in-Coding/WeatherAPI.git
cd WeatherAPI
```

### 2. Add your API key

Create a file named `api.key` in the root directory and paste your Visual Crossing API Key:

```bash
echo "your_api_key_here" > api.key
```

### 3. Start Redis with Docker

```bash
docker-compose up -d
```

This launches Redis with a persistent Docker volume.

## Usage

Run the CLI tool like this:

```bash
go run . Lahore
```

The tool will:
1. Check Redis for cached weather info
2. If not found, hit the API and cache the result
3. Print formatted weather details in terminal

## Clean up

To stop and remove the Redis container:

```bash
docker-compose down -v
```

## Project Structure

```
├── api.key              # Your API key (not committed)
├── docker-compose.yml   # Redis container setup
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
├── main.go              # Entry point
└── utils/               # Helper utilities
    ├── cache.go         # Redis caching logic
    └── jsonParse.go     # JSON parsing logic
```


## Acknowledgments

- [Visual Crossing](https://www.visualcrossing.com/) for the weather API
- [Redis](https://redis.io/) for blazing-fast caching
- [Docker](https://www.docker.com/) for containerization

## Project URL

[Roadmap.sh](https://roadmap.sh/projects/weather-api-wrapper-service)
