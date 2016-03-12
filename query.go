package webque

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"reflect"
	"unicode"

	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
	"github.com/jackc/pgx"
)

// package errors
var (
	ErrInvalidPointer = errors.New("attempt to load into an invalid pointer")
)

var (
	typeValuer = reflect.TypeOf((*driver.Valuer)(nil)).Elem()
)

func camelCaseToSnakeCase(name string) string {
	buf := new(bytes.Buffer)

	runes := []rune(name)

	for i := 0; i < len(runes); i++ {
		buf.WriteRune(unicode.ToLower(runes[i]))
		if i != len(runes)-1 && unicode.IsUpper(runes[i+1]) &&
			(unicode.IsLower(runes[i]) || unicode.IsDigit(runes[i]) ||
				(i != len(runes)-2 && unicode.IsLower(runes[i+2]))) {
			buf.WriteRune('_')
		}
	}

	return buf.String()
}

func structValue(m map[string]reflect.Value, value reflect.Value) {
	if value.Type().Implements(typeValuer) {
		return
	}
	switch value.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
			return
		}
		structValue(m, value.Elem())
	case reflect.Struct:
		t := value.Type()
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.PkgPath != "" {
				// unexported
				continue
			}
			tag := field.Tag.Get("db")
			if tag == "-" {
				// ignore
				continue
			}
			if tag == "" {
				// no tag, but we can record the field name
				tag = camelCaseToSnakeCase(field.Name)
			}
			fieldValue := value.Field(i)
			if _, ok := m[tag]; !ok {
				m[tag] = fieldValue
			}
			structValue(m, fieldValue)
		}
	}
}

func structMap(value reflect.Value) map[string]reflect.Value {
	m := make(map[string]reflect.Value)
	structValue(m, value)
	return m
}

// Load loads any value from sql.Rows
func Load(rows *pgx.Rows, value interface{}) (int, error) {
	defer rows.Close()

	var column []string
	for _, c := range rows.FieldDescriptions() {
		column = append(column, c.Name)
	}

	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return 0, ErrInvalidPointer
	}
	v = v.Elem()
	isSlice := v.Kind() == reflect.Slice && v.Type().Elem().Kind() != reflect.Uint8
	count := 0
	for rows.Next() {
		var elem reflect.Value
		if isSlice {
			elem = reflect.New(v.Type().Elem()).Elem()
		} else {
			elem = v
		}
		ptr, err := findPtr(column, elem)
		if err != nil {
			return 0, err
		}
		err = rows.Scan(ptr...)
		if err != nil {
			return 0, err
		}
		count++
		if isSlice {
			v.Set(reflect.Append(v, elem))
		} else {
			break
		}
	}
	return count, nil
}

var (
	dummyDest   interface{}
	typeScanner = reflect.TypeOf((*pgx.Scanner)(nil)).Elem()
)

func findPtr(column []string, value reflect.Value) ([]interface{}, error) {
	if value.Addr().Type().Implements(typeScanner) {
		return []interface{}{value.Addr().Interface()}, nil
	}
	switch value.Kind() {
	case reflect.Struct:
		var ptr []interface{}
		m := structMap(value)
		for _, key := range column {
			if val, ok := m[key]; ok {
				ptr = append(ptr, val.Addr().Interface())
			} else {
				ptr = append(ptr, &dummyDest)
			}
		}
		return ptr, nil
	case reflect.Ptr:
		if value.IsNil() {
			value.Set(reflect.New(value.Type().Elem()))
		}
		return findPtr(column, value.Elem())
	}
	return []interface{}{value.Addr().Interface()}, nil
}

// ToSelectSQL create select sql string
func ToSelectSQL(stmt *dbr.SelectStmt) (string, error) {
	builder := &dbr.SelectBuilder{
		Dialect:    dialect.PostgreSQL,
		SelectStmt: stmt,
	}
	sql, value := builder.ToSql()
	query, err := dbr.InterpolateForDialect(sql, value, dialect.PostgreSQL)
	if err != nil {
		return "", err
	}
	return query, nil
}

// ToInsertSQL create insert sql string
func ToInsertSQL(stmt *dbr.InsertStmt) (string, error) {
	builder := &dbr.InsertBuilder{
		Dialect:    dialect.PostgreSQL,
		InsertStmt: stmt,
	}
	sql, value := builder.ToSql()
	query, err := dbr.InterpolateForDialect(sql, value, dialect.PostgreSQL)
	if err != nil {
		return "", err
	}
	return query, nil
}

// ToUpdateSQL create update sql string
func ToUpdateSQL(stmt *dbr.UpdateStmt) (string, error) {
	builder := &dbr.UpdateBuilder{
		Dialect:    dialect.PostgreSQL,
		UpdateStmt: stmt,
		LimitCount: -1,
	}
	sql, value := builder.ToSql()
	query, err := dbr.InterpolateForDialect(sql, value, dialect.PostgreSQL)
	if err != nil {
		return "", err
	}
	return query, nil
}
