# Url Shortener

Url shortener API made with Go, Gin and Redis. Deploy to [Railway](https://railway.app) using [Redis Cloud](https://app.redislabs.com).

## Run

This project use default enviroment variables, redis addr is '127.0.0.1:6379' and other values are empty.
Unless other variables are found. Just exec:

```go
go run .
```

## Endpoints

- GET /api/all: Get all ids in database
- GET /api/:id: Get all ids in database
- POST /api/:url: Get all ids in database

## Other

This project also have a gRPC version [here](https://github.com/Patrick564/url-shortener-backend)
