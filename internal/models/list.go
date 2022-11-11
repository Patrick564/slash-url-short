package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UrlModel struct {
	ShortUrl    string `json:"name" db:"name"`
	OriginalUrl int64  `json:"age" db:"age"`
}

type UrlList []UrlModel

// @todo change func name to AllUrls()
func FetchAllUrls(ctx context.Context, db *pgxpool.Pool) error {
	rows, err := db.Query(ctx, "select name, age from users")
	if err != nil {
		return err
	}
	defer rows.Close()

	urls := make(UrlList, 0)
	for rows.Next() {
		var u UrlModel
		err = rows.Scan(&u.ShortUrl, &u.OriginalUrl)
		if err != nil {
			return err
		}
		urls = append(urls, u)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	fmt.Printf("list: %+v\n", urls)
	return nil
}
