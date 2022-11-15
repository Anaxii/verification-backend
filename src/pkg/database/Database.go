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

func RefreshQueue() {
	var _queue []global.VerificationRequest

	db, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM requests")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var walletAddress string
		var email string
		var status string
		var message string
		var account string
		var hashed_message string
		var r string
		var s string
		var v string
		var sig string
		row.Scan(&walletAddress, &email, &status, &message, &account, &hashed_message, &r, &s, &v, &sig)
		_queue = append(_queue, global.VerificationRequest{WalletAddress: walletAddress, Email: email, Signature: global.SignatureStruct{Message: message, Account: account, SignatureData: global.SignatureData{HashedMessage: hashed_message, R: r, S: s, V: v, Sig: sig}}})
	}
	global.Queue = _queue
}