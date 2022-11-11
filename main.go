package main

import (
	"context"
	"log"
	"os"

	"github.com/Patrick564/url-shortener-backend/internal/database"
)

const databaseUrl string = "postgres://golang:12345@localhost:5432/conn_test"

func main() {
	// r := api.SetupRouter()
	ctx := context.Background()
	d, err := database.OpenDatabaseConn(ctx, databaseUrl)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer d.Close()

	// allurls.FetchAllUrls(ctx, d)

	// r.Run()
}
