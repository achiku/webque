package webque

import "testing"

func TestNewProxyDB(t *testing.T) {
	db, err := NewProxyDB("postgresql://localhost/webque_proxy")
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

func TestNewBackendDB(t *testing.T) {
	db, err := NewBackendDB("postgresql://localhost/webque_backend")
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
