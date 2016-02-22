package dbrx

import (
	"log"

	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
)

// SelectStmt struct
type SelectStmt struct {
	*dbr.SelectBuilder
}

// Echo test
func (s *SelectStmt) Echo() {
	log.Println("aaa")
}

// ToSQL returns sql statement
func (s *SelectStmt) ToSQL() (string, error) {
	sql, args := s.ToSql()
	query, err := dbr.InterpolateForDialect(
		sql, args, s.Dialect)
	if err != nil {
		return "", err
	}
	return query, nil
}

// Select creates a dbr.SelectBuilder
func Select(column ...interface{}) *SelectStmt {
	stmt := &dbr.SelectStmt{
		Column:      column,
		LimitCount:  -1,
		OffsetCount: -1,
	}
	builder := &dbr.SelectBuilder{
		Dialect:    dialect.PostgreSQL,
		SelectStmt: stmt,
	}
	return &SelectStmt{builder}
}
