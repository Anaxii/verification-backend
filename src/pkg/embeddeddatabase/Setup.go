package embeddeddatabase

import (
	"database/sql"
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
)

func SetupDatabase() {
	if _, err := os.Stat("./sqlite-database.db"); errors.Is(err, os.ErrNotExist) {
		log.Println("Creating sqlite-edatabase.db...")
		file, err := os.Create("sqlite-database.db") // Create SQLite file
		if err != nil {
			log.WithFields(log.Fields{"error": err.Error(), "file": "Database:DeleteRequest"}).Fatal("Could not create database")
		}
		file.Close()
		log.Println("sqlite-database.db created")

		db, err := sql.Open("sqlite3", "./sqlite-database.db")
		if err != nil {
			log.WithFields(log.Fields{"error": err.Error(), "file": "Database:DeleteRequest"}).Fatal("Could not open database")
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
			log.WithFields(log.Fields{"error": err.Error(), "file": "Database:DeleteRequest"}).Fatal("Failed to prepare table creation")
		}
		_, err = statement.Exec()
		if err != nil {
			log.WithFields(log.Fields{"error": err.Error(), "file": "Database:DeleteRequest"}).Fatal("Failed to execute table creation")
		}

		log.Println("Created table requests")

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
		"subaccount_message",
		"subaccount_account" TEXT,
		"subaccount_hashed_message" TEXT,
		"subaccount_r" TEXT,
		"subaccount_s" TEXT,
		"subaccount_v" TEXT,
		"subaccount_sig" TEXT
	  );`)
		if err != nil {
			log.WithFields(log.Fields{"error": err.Error(), "file": "Database:DeleteRequest"}).Fatal("Failed to prepare table creation")
		}

		_, err = statement.Exec()
		if err != nil {
			log.WithFields(log.Fields{"error": err.Error(), "file": "Database:DeleteRequest"}).Fatal("Failed to execute table creation")
		}

		log.Println("Created table subaccount_requests")

	}
}
