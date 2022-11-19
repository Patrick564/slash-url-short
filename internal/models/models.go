package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Url struct {
	ShortUrl    string `json:"short_url" db:"short_url"`
	OriginalUrl string `json:"original_url" db:"original_url"`
}

type UrlModel struct {
	DB  *pgxpool.Pool
	Ctx context.Context
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

func (u UrlModel) GetByID() {}

func (u UrlModel) Add(id string, url string) (Url, error) {
	_, err := u.DB.Exec(u.Ctx, "INSERT INTO mock_values(short_url, original_url) VALUES ($1, $2)", id, url)
	if err != nil {
		return Url{}, err
	}

	return Url{ShortUrl: id, OriginalUrl: url}, nil
}

func OpenDatabaseConn(ctx context.Context, databaseUrl string) (UrlModel, error) {
	dbpool, err := pgxpool.New(ctx, databaseUrl)
	if err != nil {
		return UrlModel{}, err
	}

	return UrlModel{DB: dbpool, Ctx: ctx}, dbpool.Ping(ctx)
}
