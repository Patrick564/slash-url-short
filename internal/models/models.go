package models

import (
	"context"
	"time"

	"github.com/Patrick564/url-shortener-backend/utils"
	"github.com/go-redis/redis/v8"
)

type UrlModel struct {
	DB  *redis.Client
	Ctx context.Context
}

func (u UrlModel) Close() {
	u.DB.Close()
}

func (u UrlModel) All() ([]string, error) {
	k, _, err := u.DB.ScanType(u.Ctx, 0, "short:id:*", 0, "string").Result()
	if err != nil {
		return nil, err
	}

	return k, nil
}

func (u UrlModel) Add(sid string, url string) (string, error) {
	res, err := u.DB.Get(u.Ctx, url).Result()
	if err == nil {
		return res, utils.ErrUrlExists
	}

	p := u.DB.Pipeline()

	p.Set(u.Ctx, sid, url, time.Second*0)
	p.Set(u.Ctx, url, sid, time.Second*0)

	_, err = p.Exec(u.Ctx)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (u UrlModel) GoTo(id string) (string, error) {
	s, err := u.DB.Get(u.Ctx, id).Result()
	if err != nil {
		return "", err
	}

	return s, nil
}

func OpenDatabaseConn(ctx context.Context, redisUser, redisAddr, redisPwd string) (UrlModel, error) {
	rdb := redis.NewClient(&redis.Options{
		Username: redisUser,
		Addr:     redisAddr,
		DB:       0,
		Password: redisPwd,
	})

	return UrlModel{DB: rdb, Ctx: ctx}, nil
}
