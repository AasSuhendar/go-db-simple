package db

import (
	"database/sql"
	"time"
)

// GetConnection func for get connection SQL MariaDB
func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/go-db?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
