package webque

import (
	"github.com/gocraft/dbr"
	"github.com/jackc/pgx"
)

// LoadRequestService create load request
func LoadRequestService(tx *pgx.Tx, req LoadRequestRequest) ([]LoadRequestModel, error) {
	var reqs []LoadRequestModel
	stmt := dbr.Select("id", "account_id", "amount", "completed", "created_at", "updated_at").
		From("load_request").
		Limit(10)
	query, err := ToSelectSQL(stmt)
	if err != nil {
		return reqs, err
	}
	rows, err := tx.Query(query)
	if err != nil {
		return reqs, err
	}
	Load(rows, reqs)
	return reqs, nil
}
