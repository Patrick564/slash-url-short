package main

import (
	"context"
	"log"
	"os"

	"github.com/Patrick564/url-shortener-backend/internal/database"
	allurls "github.com/Patrick564/url-shortener-backend/models/all_urls"
)

func main() {
	// r := api.SetupRouter()
	ctx := context.Background()

	p, err := database.New(ctx)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer p.Close()

	allurls.FetchAllUrls(ctx, p.DB)

	// r.Run()
}
