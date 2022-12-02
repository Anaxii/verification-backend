package embeddeddatabase

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"puffinverificationbackend/src/pkg/global"
)

func InsertNewSubAccountRequest(data global.SubAccountRequest, id primitive.ObjectID) error {
	db, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Database:InsertNewSubAccountRequest"}).Fatal("Could not open database")
	}
	defer db.Close()

	statement, err := db.Prepare(
		`INSERT INTO subaccount_requests(
                     parent_wallet_address, subaccount_wallet_address, id, status, 
                     parent_hashed_message, parent_r, parent_s, parent_v, parent_sig, 
                     parent_message, parent_account, subaccount_hashed_message, subaccount_r, 
                     subaccount_s, subaccount_v, subaccount_sig, subaccount_message, subaccount_account) 
                     VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Database:InsertNewSubAccountRequest"}).Error("Could not prepare for insert")
		return err
	}
	_, err = statement.Exec(
		data.ParentAddress,
		data.SubAccountAddress,
		id.Hex(),
		"pending",
		data.ParentSignature.SignatureData.HashedMessage,
		data.ParentSignature.SignatureData.R,
		data.ParentSignature.SignatureData.S,
		data.ParentSignature.SignatureData.V,
		data.ParentSignature.SignatureData.Sig,
		data.ParentSignature.Message,
		data.ParentSignature.Account,
		data.SubAccountSignature.SignatureData.HashedMessage,
		data.SubAccountSignature.SignatureData.R,
		data.SubAccountSignature.SignatureData.S,
		data.SubAccountSignature.SignatureData.V,
		data.SubAccountSignature.SignatureData.Sig,
		data.SubAccountSignature.Message,
		data.SubAccountSignature.Account,
	)
	if err != nil {
		return err
	}
	return nil
}

func InsertNewRequest(data global.VerificationRequest, id primitive.ObjectID) error {
	db, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Database:InsertNewRequest"}).Fatal("Could not open database")
	}
	defer db.Close()

	statement, err := db.Prepare(
		`INSERT INTO requests(wallet_address, id, status, email, hashed_message, r, s, v, sig, message, account) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Database:InsertNewRequest"}).Error("Could not prepare for insert")
		return err
	}
	_, err = statement.Exec(
		data.WalletAddress,
		id.Hex(),
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

func DeleteRequest(table string, walletAddress string, name string) error {
	db, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Database:DeleteRequest"}).Fatal("Could not open database")
	}
	defer db.Close()

	statement, err := db.Prepare(
		`DELETE FROM ` + table + ` WHERE ` + name + ` = ?`,
	)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Database:DeleteRequest"}).Error("Could not prepare for delete")
		return err
	}
	_, err = statement.Exec(
		walletAddress,
	)
	if err != nil {
		return err
	}
	return nil
}

func RefreshQueue() {
	var _queue []global.VerificationRequest
	var _subaccountQueue []global.SubAccountRequest

	db, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Database:RefreshQueue"}).Fatal("Could not open database")
		return
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM requests")
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Database:RefreshQueue"}).Warn("Failed to query db table requests")
		return
	}
	defer row.Close()
	for row.Next() {
		var walletAddress string
		var email string
		var id string
		var status string
		var message string
		var account string
		var hashed_message string
		var r string
		var s string
		var v string
		var sig string

		err = row.Scan(&walletAddress, &id, &email, &status, &message, &account, &hashed_message, &r, &s, &v, &sig)
		if err != nil {
			continue
		}
		_queue = append(_queue, global.VerificationRequest{WalletAddress: walletAddress, ID: id, Status: status, Email: email, Signature: global.SignatureStruct{Message: message, Account: account, SignatureData: global.SignatureData{HashedMessage: hashed_message, R: r, S: s, V: v, Sig: sig}}})
	}
	global.Queue = _queue

	row, err = db.Query("SELECT * FROM subaccount_requests")
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Database:RefreshQueue"}).Warn("Failed to query db table subaccount_requests")
		return
	}
	defer row.Close()
	for row.Next() {
		var parentWalletAddress string
		var subaccountWalletAddress string
		var id string
		var status string

		var parentMessage string
		var parentAccount string
		var parentHashedMessage string
		var parentR string
		var parentS string
		var parentV string
		var parentSig string

		var subaccountMessage string
		var subaccountAccount string
		var subaccountHashedMessage string
		var subaccountR string
		var subaccountS string
		var subaccountV string
		var subaccountSig string

		err = row.Scan(&parentWalletAddress, &subaccountWalletAddress, &id, &status,
			&parentMessage, &parentAccount, &parentHashedMessage, &parentR, &parentS, &parentV, &parentSig,
			&subaccountMessage, &subaccountAccount, &subaccountHashedMessage, &subaccountR, &subaccountS, &subaccountV, &subaccountSig)
		if err != nil {
			continue
		}
		_subaccountQueue = append(_subaccountQueue, global.SubAccountRequest{
			ParentAddress: parentWalletAddress, SubAccountAddress: subaccountWalletAddress, ID: id,
			ParentSignature: global.SignatureStruct{
				Message: parentMessage, Account: parentAccount, SignatureData: global.SignatureData{HashedMessage: parentHashedMessage, R: parentR, S: parentS, V: parentV, Sig: parentSig},
			},
			SubAccountSignature: global.SignatureStruct{
				Message: subaccountMessage, Account: subaccountWalletAddress, SignatureData: global.SignatureData{HashedMessage: subaccountHashedMessage, R: subaccountR, S: subaccountS, V: subaccountV, Sig: subaccountSig},
			}})
	}
	global.SubAccountQueue = _subaccountQueue
}
