package main

import "github.com/Patrick564/url-shortener-backend/api"

func main() {
	r := api.SetupRouter()
	r.Run()
}
