package webque

import (
	"testing"

	"github.com/gocraft/dbr"
)

func TestNewDB(t *testing.T) {
	db, err := NewDB("postgresql://localhost/webque_proxy")
	if err != nil {
		t.Fatal(err)
	}
	var n int32
	err = db.QueryRow("select 100").Scan(&n)
	if err != nil {
		t.Error(err)
	}
	t.Logf("got %d", n)
}

func TestToSelectSQLNoPlaceholder(t *testing.T) {
	stmt := dbr.Select("id", "name").
		From("account").
		OrderDesc("id")
	query, err := ToSelectSQL(stmt)
	if err != nil {
		t.Error(err)
	}

	targetSQL := "SELECT id, name FROM account ORDER BY id DESC"
	if query != targetSQL {
		t.Errorf("want %s, got %s", targetSQL, query)
	}
	t.Log(query)
}

func TestToSelectSQLWithPlaceholder(t *testing.T) {
	stmt := dbr.Select("id", "name").
		From("account").
		Where("id = ?", 1)
	query, err := ToSelectSQL(stmt)
	if err != nil {
		t.Error(err)
	}

	targetSQL := "SELECT id, name FROM account WHERE (id = 1)"
	if query != targetSQL {
		t.Errorf("want %s, got %s", targetSQL, query)
	}
	t.Log(query)
}

func TestToInsertSQL(t *testing.T) {
	type user struct {
		ID   int
		Name string
	}
	stmt := dbr.InsertInto("account").Columns("id", "name").
		Record(&user{ID: 1, Name: "8maki"}).
		Record(&user{ID: 2, Name: "moqada"})
	query, err := ToInsertSQL(stmt)
	if err != nil {
		t.Error(err)
	}
	t.Log(query)
}
