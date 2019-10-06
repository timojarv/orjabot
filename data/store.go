package data

import (
	"log"
	"database/sql"
	// Driver
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./orja.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := newMigration(db, "sqlite3").Do(false); err != nil {
		log.Fatal(err)
	}
}
