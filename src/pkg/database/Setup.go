package database

import (
	"database/sql"
	"errors"
	"log"
	"os"
)

func SetupDatabase() {
	if _, err := os.Stat("./sqlite-database.db"); errors.Is(err, os.ErrNotExist) {
		log.Println("Creating sqlite-database.db...")
		file, err := os.Create("sqlite-database.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("sqlite-database.db created")

		db, err := sql.Open("sqlite3", "./sqlite-database.db")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer db.Close()

		log.Println("Create tables...")
		statement, err := db.Prepare(`CREATE TABLE requests (
		"wallet_address" TEXT,		
		"email" TEXT,
		"status" TEXT,
		"message" TEXT,
		"account" TEXT,
		"hashed_message" TEXT,
		"r" TEXT,
		"s" TEXT,
		"v" TEXT,
		"sig" TEXT
	  );`)
		if err != nil {
			log.Fatal(err.Error())
		}
		_, err = statement.Exec() // Execute SQL Statements
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("Table created")
	}
}
