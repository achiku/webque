package webque

import (
	"testing"

	"github.com/jackc/pgx"
)

func createUserTestData(tx *pgx.Tx) error {
	users := []string{"achiku", "moqada", "ideyuta"}
	for _, u := range users {
		_, err := tx.Exec(`insert into account (name) values ($1)`, u)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestLoadRequestService(t *testing.T) {
	db, err := NewDB("postgresql://localhost/webque_proxy")
	if err != nil {
		t.Fatal(err)
	}

	req := LoadRequestRequest{
		AccountID: 1,
		Amount:    1000,
	}
	tx, err := db.Begin()
	defer tx.Rollback()
	if err = createUserTestData(tx); err != nil {
		t.Error(err)
	}
	resp, err := LoadRequestService(tx, req)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
