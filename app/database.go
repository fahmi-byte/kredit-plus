package app

import (
	"database/sql"
	_ "github.com/lib/pq"
	"kredit-plus/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:123@localhost/kredit_plus_db?sslmode=disable")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
