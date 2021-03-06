package webque

import (
	"time"

	"github.com/gocraft/dbr"
	"github.com/jackc/pgx"
)

// GetLoadRequestService get load request
func GetLoadRequestService(tx *pgx.Tx, req GetLoadRequestRequest) ([]LoadRequestModel, error) {
	var reqs []LoadRequestModel
	stmt := dbr.Select(
		"id",
		"account_id",
		"amount",
		"completed",
		"created_at",
		"updated_at",
	).
		From("load_request").
		Where("account_id = ?", req.AccountID).
		Limit(10)
	query, err := ToSelectSQL(stmt)
	if err != nil {
		return reqs, err
	}
	rows, err := tx.Query(query)
	if err != nil {
		return reqs, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			ID        int32
			accountID int32
			amount    int32
			completed bool
			createdAt time.Time
			updatedAt time.Time
		)
		err = rows.Scan(&ID, &accountID, &amount, &completed, &createdAt, &updatedAt)
		if err != nil {
			return reqs, err
		}
		r := LoadRequestModel{
			ID:        ID,
			AccountID: accountID,
			Amount:    amount,
			Completed: completed,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		reqs = append(reqs, r)
	}
	return reqs, nil
}

// CreateLoadRequestService creates load request
func CreateLoadRequestService(tx *pgx.Tx, req LoadRequestRequest) error {
	stmt := dbr.InsertInto("load_request").
		Columns("account_id", "amount").
		Values(req.AccountID, req.Amount)
	query, err := ToInsertSQL(stmt)
	if err != nil {
		return err
	}
	_, err = tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// UpdateLoadRequestService updates load request
func UpdateLoadRequestService(tx *pgx.Tx, req LoadRequestRequest) error {
	stmt := dbr.Update("load_request").
		Set("completed", true).
		Where("id = ?", req.RequestID)
	query, err := ToUpdateSQL(stmt)
	if err != nil {
		return err
	}
	_, err = tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
