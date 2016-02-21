package qb

import "testing"

func TestSelect(t *testing.T) {
	stmt := Select("id", "name").
		From("account").
		Where("id = $1", 1).
		OrderDir("id", false)
	sql, args := stmt.ToSql()
	t.Log(sql, args)
}
