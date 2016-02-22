package webque

import (
	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
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
	}
	pool, err := pgx.NewConnPool(poolcfg)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

// ToSelectSQL create select sql string
func ToSelectSQL(stmt *dbr.SelectStmt) (string, error) {
	builder := &dbr.SelectBuilder{
		Dialect:    dialect.PostgreSQL,
		SelectStmt: stmt,
	}
	sql, value := builder.ToSql()
	query, err := dbr.InterpolateForDialect(sql, value, dialect.PostgreSQL)
	if err != nil {
		return "", err
	}
	return query, nil
}

// ToInsertSQL create insert sql string
func ToInsertSQL(stmt *dbr.InsertStmt) (string, error) {
	builder := &dbr.InsertBuilder{
		Dialect:    dialect.PostgreSQL,
		InsertStmt: stmt,
	}
	sql, value := builder.ToSql()
	query, err := dbr.InterpolateForDialect(sql, value, dialect.PostgreSQL)
	if err != nil {
		return "", err
	}
	return query, nil
}
