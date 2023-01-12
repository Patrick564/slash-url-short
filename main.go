package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/Patrick564/url-shortener-backend/api"
	"github.com/Patrick564/url-shortener-backend/api/controllers"
	"github.com/Patrick564/url-shortener-backend/internal/models"
	_ "github.com/joho/godotenv"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("error at loading .env file: %+v", err)
	// }

	ctx := context.Background()

	redisHost := os.Getenv("REDIS_HOST")
	redisDb, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	redisPwd := os.Getenv("REDIS_PWD")

	u, err := models.OpenDatabaseConn(ctx, redisDb, redisHost, redisPwd)
	if err != nil {
		log.Fatalln(err)
	}
	defer u.Close()

	e := &controllers.Env{Urls: u}

	r := api.SetupRouter(e)
	r.Run()
}
