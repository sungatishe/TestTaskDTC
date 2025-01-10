package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var DB *sql.DB

func InitDB() {
	dsn := os.Getenv("dsn")
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic("Failed to connect db ")
	}

	err = DB.Ping()
	if err != nil {
		panic("Failed to ping db ")
	}

	if err != nil {
		panic(fmt.Sprintf("Failed to create todos table: %v", err))
	}
}
