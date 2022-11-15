package verify

import (
	"log"
	"puffinverificationbackend/src/pkg/blockchain"
	"puffinverificationbackend/src/pkg/database"
	"puffinverificationbackend/src/pkg/global"
)

func HandleRequests() {
	for {
		select {
		case <-global.Check:
			database.RefreshQueue()
			for _, v := range global.Queue {
				isValid := blockchain.VerifySignature(v.Signature.SignatureData, v.WalletAddress)

				if !isValid {
					database.DeleteRequest(v)
					continue
				}

				if !blockchain.CheckIfIsApproved(v.WalletAddress) {
					blockchain.ApproveAddress(v.WalletAddress)
					// send to mongodb

					//database.DeleteRequest(v)
				}
				log.Println(blockchain.CheckIfIsApproved(v.WalletAddress))
			}
		}
	}
}

