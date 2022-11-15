package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	DB  *pgxpool.Pool
	Ctx context.Context
}

func (d *Database) Close() {
	d.DB.Close()
}

func OpenDatabaseConn(ctx context.Context, databaseUrl string) (*Database, error) {
	dbpool, err := pgxpool.New(ctx, databaseUrl)
	if err != nil {
		return nil, err
	}

	return &Database{DB: dbpool, Ctx: ctx}, nil
}
