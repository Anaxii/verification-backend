package verify

import (
	"log"
	"puffinverificationbackend/src/pkg/blockchain"
	"puffinverificationbackend/src/pkg/database"
	"puffinverificationbackend/src/pkg/global"
	"time"
)

func minuteTicker() *time.Ticker {
	return time.NewTicker(time.Second * time.Duration(60-time.Now().Second()))
}

func HandleRequests() {
	updating := false
	go func() {
		log.Println("Started ticker")
		t := minuteTicker()
		for {
			<-t.C
			t = minuteTicker()

			if updating {
				global.CheckRequests <- true
			}
		}
	}()
	for {
		select {
		case <-global.CheckRequests:
			updating = true
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
			updating = false
		}
	}
}

