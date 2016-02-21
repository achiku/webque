package qb

import (
	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
)

// Select creates a dbr.SelectBuilder
func Select(column ...interface{}) *dbr.SelectBuilder {
	stmt := &dbr.SelectStmt{
		Column:      column,
		LimitCount:  -1,
		OffsetCount: -1,
	}
	return &dbr.SelectBuilder{
		Dialect:    dialect.PostgreSQL,
		SelectStmt: stmt,
	}
}
