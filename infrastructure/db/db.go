package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DBInit() (*sql.DB, error) {
	var err error
	DB, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
	return DB, err
}

func DBClose() {
	if DB != nil {
		DB.Close()
	}
}

func DBConn() *sql.DB {
	return DB
}
