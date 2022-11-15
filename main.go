package main

import (
	"context"
	"log"
	"os"

	"github.com/Patrick564/url-shortener-backend/api"
	"github.com/Patrick564/url-shortener-backend/internal/models"
)

const databaseUrl string = "postgres://golang:12345@localhost:5432/conn_test"

func main() {
	ctx := context.Background()
	d, err := models.OpenDatabaseConn(ctx, databaseUrl)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer d.Close()

	r := api.SetupRouter(d.DB)
	r.Run()
}
