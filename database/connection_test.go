package database

import "testing"

func TestGetConnection(t *testing.T) {

	db := getConnection()

	// as variáveis usadas no sql devem ser números apenas ($1, $2, $3, ...)
	sql := `INSERT INTO "Employee" ("name") VALUES ($1)`

	// o primeiro parâmetro é o sql e os demais parâmetros representam respectivamente as variáveis declaradas no sql
	db.Exec(sql, "Rafael S.")
	defer db.Close()
}