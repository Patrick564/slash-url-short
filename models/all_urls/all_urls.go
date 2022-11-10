package allurls

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

func FetchAllUrls(ctx context.Context, db *pgxpool.Pool) error {
	rows, err := db.Query(ctx, "select name, age from users")
	if err != nil {
		// fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		// os.Exit(1)
		return err
	}
	defer rows.Close()

	var urls UrlList
	for rows.Next() {
		var u UrlModel
		err = rows.Scan(&u.ShortUrl, &u.OriginalUrl)
		if err != nil {
			// fmt.Fprintln(os.Stderr, "aaaaaa")
			// os.Exit(1)
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
