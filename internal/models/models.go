package models

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Url struct {
	ShortUrl    string `json:"short_url" db:"short_url"`
	OriginalUrl string `json:"original_url" db:"original_url"`
}

type UrlModel struct {
	DB  *redis.Client
	Ctx context.Context
}

func (u UrlModel) Close() {
	u.DB.Close()
}

func (u UrlModel) All() ([]string, error) {
	k, _, err := u.DB.ScanType(u.Ctx, 0, "", 0, "hash").Result()
	if err != nil {
		return nil, err
	}

	return k, nil
}

func (u UrlModel) Add(sid string, rawUrl string) (string, error) {
	key := fmt.Sprintf("url:%s", sid)

	err := u.DB.HSet(u.Ctx, key, "short", sid, "original", rawUrl).Err()
	if err != nil {
		return "", err
	}

	return sid, nil
}

func (u UrlModel) GoTo(id string) (string, error) {
	key := fmt.Sprintf("url:%s", id)
	s, err := u.DB.HGet(u.Ctx, key, "original").Result()
	if err != nil {
		return "", err
	}

	return s, nil
}

func OpenDatabaseConn(ctx context.Context, redisUrl, redisPwd string) (UrlModel, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		DB:       0,
		Password: redisPwd,
	})

	return UrlModel{DB: rdb, Ctx: ctx}, nil
}
