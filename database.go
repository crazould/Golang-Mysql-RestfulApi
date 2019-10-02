package main

import (
	"database/sql"
	f "fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//DbUsername ...
const DbUsername = "root"

//DbPassword ...
const DbPassword = ""

//DbServerAddress ...
const DbServerAddress = "127.0.0.1"

//DbServerPort ...
const DbServerPort = "3306"

//DbName ...
const DbName = "mydatabase"

//DbHandler ...
type DbHandler struct {
	ConnectionString string
}

//Connect ...
func (db *DbHandler) Connect() *sql.DB {
	db.ConnectionString = f.Sprintf("%s:%s@tcp(%s:%s)/%s",
		DbUsername,
		DbPassword,
		DbServerAddress,
		DbServerPort,
		DbName,
	)

	conn, err := sql.Open("mysql", db.ConnectionString)

	if err != nil {
		log.Panic(err)
	}

	return conn
}

// Query ...
func (db *DbHandler) Query(sql string) (*sql.Rows, error) {
	conn := db.Connect()
	return conn.Query(sql)
}
