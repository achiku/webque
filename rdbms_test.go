package webque

import "testing"

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
