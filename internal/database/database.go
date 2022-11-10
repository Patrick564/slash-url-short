package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	databaseUrl string = "postgres://golang:12345@localhost:5432/conn_test"
)

// Using only pgxpool and not sql.DB interface because
// other db's compatibility is not required in this project.
type PoolConn struct {
	DB *pgxpool.Pool
}

func (p *PoolConn) StartConn(ctx context.Context) error {
	pool, err := pgxpool.New(ctx, databaseUrl)
	if err != nil {
		return err
	}

	p.DB = pool

	return nil
}

func (p *PoolConn) Close() {
	p.DB.Close()
}

func New(ctx context.Context) (*PoolConn, error) {
	p := PoolConn{}

	err := p.StartConn(ctx)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
