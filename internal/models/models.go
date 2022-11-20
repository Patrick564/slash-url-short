package models

import (
	"context"
	"net/url"

	"github.com/Patrick564/url-shortener-backend/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/teris-io/shortid"
)

type Url struct {
	ShortUrl    string `json:"short_url" db:"short_url"`
	OriginalUrl string `json:"original_url" db:"original_url"`
}

type UrlModel struct {
	DB  *pgxpool.Pool
	Ctx context.Context
	SID *shortid.Shortid
}

func (u UrlModel) Close() {
	u.DB.Close()
}

func (u UrlModel) All() ([]Url, error) {
	rows, err := u.DB.Query(u.Ctx, "SELECT short_url, original_url FROM mock_values")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []Url
	for rows.Next() {
		var u Url
		err = rows.Scan(&u.ShortUrl, &u.OriginalUrl)
		if err != nil {
			return nil, err
		}
		urls = append(urls, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

func (u UrlModel) Add(rawUrl string) (Url, error) {
	p, err := url.ParseRequestURI(rawUrl)
	if err != nil || p.Scheme == "" || p.Host == "" {
		return Url{}, utils.ErrInvalidUrl
	}

	id, err := u.SID.Generate()
	if err != nil {
		return Url{}, err
	}

	_, err = u.DB.Exec(u.Ctx, "INSERT INTO mock_values(short_url, original_url) VALUES ($1, $2)", id, p)
	if err != nil {
		return Url{}, err
	}

	return Url{ShortUrl: id, OriginalUrl: p.String()}, nil
}

func OpenDatabaseConn(ctx context.Context, databaseUrl string) (UrlModel, error) {
	dbpool, err := pgxpool.New(ctx, databaseUrl)
	if err != nil {
		return UrlModel{}, err
	}

	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		return UrlModel{}, err
	}

	return UrlModel{DB: dbpool, Ctx: ctx, SID: sid}, dbpool.Ping(ctx)
}
