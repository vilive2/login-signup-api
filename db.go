package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func GetDB() *sql.DB {

	var err error
	if db == nil {
		db, err = sql.Open("sqlite3", "db.db")
		if err != nil {
			log.Println("error connecting db:", err)
		}

		if err == nil {
			_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS users(
				username VARCHAR(32) PRIMARY KEY,
				name VARCHAR(64),
				email VARCHAR(64) NOT NULL UNIQUE,
				password VARCHAR(64) NOT NULL
			)
			`)

			if err != nil {
				log.Println("error creating table:", err)
			}
		}
	}

	return db
}
