package webque

import (
	"github.com/bgentry/que-go"
	"github.com/jackc/pgx"
)

// NewProxyDB create DB
func NewProxyDB(dbURI string) (*pgx.ConnPool, error) {
	pgxcfg, err := pgx.ParseURI(dbURI)
	if err != nil {
		return nil, err
	}
	poolcfg := pgx.ConnPoolConfig{
		ConnConfig:   pgxcfg,
		AfterConnect: que.PrepareStatements,
	}
	pool, err := pgx.NewConnPool(poolcfg)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

// NewBackendDB create DB
func NewBackendDB(dbURI string) (*pgx.ConnPool, error) {
	pgxcfg, err := pgx.ParseURI(dbURI)
	if err != nil {
		return nil, err
	}
	poolcfg := pgx.ConnPoolConfig{
		ConnConfig: pgxcfg,
	}
	pool, err := pgx.NewConnPool(poolcfg)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
