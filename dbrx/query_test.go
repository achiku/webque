package dbrx

import "testing"

func TestSelect(t *testing.T) {
	stmt := Select("id", "name")
	stmt.Echo()
}
