package embeddeddatabase

import (
	"database/sql"
	"errors"
	"log"
	"os"
)

func SetupDatabase() {
	if _, err := os.Stat("./sqlite-database.db"); errors.Is(err, os.ErrNotExist) {
		log.Println("Creating sqlite-edatabase.db...")
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
		"id" TEXT,	
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

		statement, err = db.Prepare(`CREATE TABLE subaccount_requests (
		"parent_wallet_address" TEXT,
		"subaccount_wallet_address" TEXT,	
		"id" TEXT,	
		"status" TEXT,
		"parent_message" TEXT,
		"parent_account" TEXT,
		"parent_hashed_message" TEXT,
		"parent_r" TEXT,
		"parent_s" TEXT,
		"parent_v" TEXT,
		"parent_sig" TEXT,
		"subaccount_account" TEXT,
		"subaccount_hashed_message" TEXT,
		"subaccount_r" TEXT,
		"subaccount_s" TEXT,
		"subaccount_v" TEXT,
		"subaccount_sig" TEXT
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
