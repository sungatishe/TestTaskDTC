package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(dsn string) {
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic("Failed to connect database ")
	}

	err = DB.Ping()
	if err != nil {
		panic("Failed to ping database ")
	}

	if err != nil {
		panic(fmt.Sprintf("Failed to create todos table: %v", err))
	}
}
