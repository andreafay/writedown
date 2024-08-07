package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func openDB() error {
	db, err := sql.Open("sqlite3", "./sqlite3.db")
	if err != nil {
		return err
	}

	DB = db
	log.Println("Database opened successfully")
	return nil
}

func closeDB() error {
	if err := DB.Close(); err != nil {
		return err
	}
	log.Println("Database closed successfully")
	return nil
}

func setupDB() error {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS notes (id INTEGER NOT NULL PRIMARY KEY, title TEXT, category TEXT);`)
	if err != nil {
		return err
	}
	log.Println("Table created or already exists")
	return nil
}
