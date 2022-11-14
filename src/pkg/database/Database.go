package database

import (
	"database/sql"
	"log"
	"puffinverificationbackend/src/pkg/global"
)

func InsertNewRequest(data global.VerificationRequest) error {
	db, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	statement, err := db.Prepare(
		`INSERT INTO requests(wallet_address, status, email, hashed_message, r, s, v, sig, message, account) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = statement.Exec(
		data.WalletAddress,
		"pending",
		data.Email,
		data.Signature.SignatureData.HashedMessage,
		data.Signature.SignatureData.R,
		data.Signature.SignatureData.S,
		data.Signature.SignatureData.V,
		data.Signature.SignatureData.Sig,
		data.Signature.Message,
		data.Signature.Account,
	)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRequest(data global.VerificationRequest) error {
	db, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	statement, err := db.Prepare(
		`DELETE FROM requests WHERE wallet_address = ?`,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = statement.Exec(
		data.WalletAddress,
	)
	if err != nil {
		return err
	}
	return nil
}