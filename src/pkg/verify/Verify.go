package verify

import (
	"puffinverificationbackend/src/pkg/blockchain"
	"puffinverificationbackend/src/pkg/embeddeddatabase"
	"puffinverificationbackend/src/pkg/externaldatabase"
	"puffinverificationbackend/src/pkg/global"
	"time"
)

func minuteTicker() *time.Ticker {
	return time.NewTicker(time.Second * time.Duration(60-time.Now().Second()))
}

func HandleRequests() {
	updating := false
	go func() {
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
			embeddeddatabase.RefreshQueue()
			for _, v := range global.Queue {
				isValid := blockchain.VerifySignature(v.Signature.SignatureData, v.WalletAddress)
				if !isValid {
					embeddeddatabase.DeleteRequest(v)
					externaldatabase.DenyRequest(v, "invalid signature")
					continue
				}

				if !blockchain.CheckIfIsApproved(v.WalletAddress) {
					err := blockchain.ApproveAddress(v.WalletAddress)
					if err != nil {
						err = externaldatabase.DenyRequest(v, "error approving wallet on mainnet")
						if err != nil {
							embeddeddatabase.DeleteRequest(v)
						}
					}
					err = blockchain.EnableOnPuffin(v.WalletAddress)
					if err != nil {
						err = externaldatabase.DenyRequest(v, "error enabling wallet on puffin")
						if err != nil {
							embeddeddatabase.DeleteRequest(v)
						}
					}
					err = externaldatabase.ApproveRequest(v)
					if err == nil {
						embeddeddatabase.DeleteRequest(v)
					}
				} else {
					err := externaldatabase.DenyRequest(v, "wallet already has kyc")
					if err != nil {
						embeddeddatabase.DeleteRequest(v)
					}
				}
			}
			updating = false
		}
	}
}

