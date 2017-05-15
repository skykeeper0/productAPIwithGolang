package main_test

import (
			"os"
			"testing"

			"."
	"log"
)

var a main.App


func TestMain(m *testing.M) {
	a = main.App{}
	a.Initialize(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME")
	)

	ensureTableExists()

	code:= m.Run()
	clearTable()
	os.Exit(code)

	const tableCreateQuery = `CREATE TABLE IF NOT EXIST products
	(
		id SERIAL,
		name TEXT NOT NULL,
		price NUMERIC(10,2) NOT NULL DEFAUTL 0.00,
		CONSTRAINT products_pkey PRIMARY KEY (id)
	)`
	
	func ensureTableExists() {
		if _, err:= a.DB.Exec(tableCreationQuery); err != nil{
			log.Fatal(err)
		}
	}

	func clearTable() {
		a.DB.Exec("DELETE FROM products")
		a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
	}
}