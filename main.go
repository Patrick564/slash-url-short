package main

import (
	"context"
	"log"
	"os"

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

	redisAddr := os.Getenv("REDIS_HOST")
	redisUser := os.Getenv("REDIS_USER")
	redisPwd := os.Getenv("REDIS_PASSWORD")

	u, err := models.OpenDatabaseConn(ctx, redisUser, redisAddr, redisPwd)
	if err != nil {
		log.Fatalln(err)
	}
	defer u.Close()

	e := &controllers.Env{Urls: u}

	r := api.SetupRouter(e)
	r.Run()
}
