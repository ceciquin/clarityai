package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() error {
	connStr := "postgres://clarityai:challenge@host:5432/clarityaidb?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *sql.DB {
	return db
}

func Close() error {
	return db.Close()
}

func Init() error {
	err := Connect()
	if err != nil {
		return err
	}

	// Run database migrations and seed data here

	return nil
}
