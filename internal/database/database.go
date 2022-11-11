package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	db *pgxpool.Pool
}

func (d *Database) Close() {
	d.db.Close()
}

// Just test add a query method
func (d *Database) Query() {}

func OpenDatabaseConn(ctx context.Context, databaseUrl string) (*Database, error) {
	dbpool, err := pgxpool.New(ctx, databaseUrl)
	if err != nil {
		return nil, err
	}

	return &Database{db: dbpool}, nil
}
