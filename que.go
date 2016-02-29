package webque

import (
	"log"

	"github.com/bgentry/que-go"
	"github.com/jackc/pgx"
)

// NewQueClient create que client
func NewQueClient(dbURI string) (*que.Client, error) {
	pgxcfg, err := pgx.ParseURI(dbURI)
	if err != nil {
		log.Fatal(err)
	}
	var qc *que.Client
	pgxpool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:   pgxcfg,
		AfterConnect: que.PrepareStatements,
	})
	if err != nil {
		log.Fatal(err)
		return qc, err
	}

	log.Println("create que client")
	qc = que.NewClient(pgxpool)
	return qc, nil
}
