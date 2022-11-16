package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Url struct {
	ShortUrl    string `json:"name" db:"short_url"`
	OriginalUrl string `json:"age" db:"original_url"`
}

type Urls []Url

type UrlModel struct {
	DB  *pgxpool.Pool
	Ctx context.Context
}

func (u *UrlModel) All() (Urls, error) {
	rows, err := u.DB.Query(u.Ctx, "SELECT short_url, original_url FROM mock_values")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	urls := make(Urls, 0)
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

func (u *Urls) GetByID() {}
