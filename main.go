package main

import (
	"context"
	"log"

	"github.com/Patrick564/url-shortener-backend/api"
	"github.com/Patrick564/url-shortener-backend/api/controllers"
	"github.com/Patrick564/url-shortener-backend/internal/models"
)

const databaseUrl string = "postgres://golang:12345@localhost:5432/conn_test"

func main() {
	ctx := context.Background()

	u, err := models.OpenDatabaseConn(ctx, databaseUrl)
	if err != nil {
		log.Fatalln(err)
	}
	// defer u.Close()

	e := &controllers.Env{Urls: u}

	r := api.SetupRouter(e)
	r.Run()
}
