package webque

import (
	"testing"

	"github.com/jackc/pgx"
)

func createUserTestData(tx *pgx.Tx) error {
	users := []string{"achiku", "moqada", "ideyuta"}
	for _, u := range users {
		_, err := tx.Exec(`INSERT INTO account (name) VALUES ($1)`, u)
		if err != nil {
			return err
		}
	}
	_, err := tx.Exec(`
	INSERT INTO load_request (account_id, amount, completed) 
	VALUES (1, 1000, false), (1, 4000, false) , (1, 11000, false)
	`)
	if err != nil {
		return err
	}
	return nil
}

func TestGetLoadRequestService(t *testing.T) {
	db, err := NewDB("postgresql://localhost/webque_proxy")
	if err != nil {
		t.Fatal(err)
	}

	req := GetLoadRequestRequest{
		AccountID: 1,
	}
	tx, err := db.Begin()
	defer tx.Rollback()
	if err = createUserTestData(tx); err != nil {
		t.Error(err)
	}
	resp, err := GetLoadRequestService(tx, req)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestCreateLoadRequestService(t *testing.T) {
	db, err := NewDB("postgresql://localhost/webque_proxy")
	if err != nil {
		t.Fatal(err)
	}

	req := LoadRequestRequest{
		AccountID: 1,
		Amount:    10000,
	}
	tx, err := db.Begin()
	defer tx.Rollback()
	if err = createUserTestData(tx); err != nil {
		t.Error(err)
	}

	if err := CreateLoadRequestService(tx, req); err != nil {
		t.Error(err)
	}
}
