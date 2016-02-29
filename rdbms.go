package webque

import (
	"github.com/bgentry/que-go"
	"github.com/jackc/pgx"
)

// NewDB create DB
func NewDB(dbURI string) (*pgx.ConnPool, error) {
	pgxcfg, err := pgx.ParseURI(dbURI)
	if err != nil {
		return nil, err
	}
	poolcfg := pgx.ConnPoolConfig{
		ConnConfig:     pgxcfg,
		MaxConnections: 5,
		AfterConnect:   que.PrepareStatements,
	}
	pool, err := pgx.NewConnPool(poolcfg)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
