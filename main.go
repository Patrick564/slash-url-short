package main

import (
	"context"
	"log"

	"github.com/Patrick564/url-shortener-backend/api"
	"github.com/Patrick564/url-shortener-backend/api/controllers"
	"github.com/Patrick564/url-shortener-backend/internal/models"
	"github.com/Patrick564/url-shortener-backend/utils"
)

func main() {
	ctx := context.Background()
	s := utils.LoadSecrets()

	u, err := models.OpenDatabaseConn(ctx, s.RedisUser, s.RedisAddr, s.RedisPwd)
	if err != nil {
		log.Fatalln(err)
	}
	defer u.Close()

	e := &controllers.Env{Urls: u}

	r := api.SetupRouter(e)
	r.Run()
}
