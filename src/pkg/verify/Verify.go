package verify

import (
	"database/sql"
	"log"
	"puffinverificationbackend/src/pkg/blockchain"
	"puffinverificationbackend/src/pkg/database"
	"puffinverificationbackend/src/pkg/global"
)

func HandleRequests() {
	for {
		select {
		case <-global.Check:
			refreshQueue()
			for _, v := range global.Queue {
				isValid := blockchain.VerifySignature(v.Signature.SignatureData, v.WalletAddress)

				if !isValid {
					database.DeleteRequest(v)
					continue
				}
				log.Println(isValid)

				if !blockchain.CheckIfIsApproved(v.WalletAddress) {
					blockchain.ApproveAddress(v.WalletAddress)
				}
				log.Println(blockchain.CheckIfIsApproved(v.WalletAddress))
			}
		}
	}
}

func refreshQueue() {
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

