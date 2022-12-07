package verify

import (
	log "github.com/sirupsen/logrus"
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
			log.WithFields(log.Fields{"file": "Verify:HandleRequests", "kyc_queue_size": len(global.Queue)}).Info("Checking KYC Queue")
			go global.Log(map[string]interface{}{"status": "checking queue", "message": "verifying kyc requests", "queue_size": len(global.Queue)})

			for _, v := range global.Queue {
				isValid := blockchain.VerifySignature(v.Signature.SignatureData, v.WalletAddress)
				if !isValid {
					embeddeddatabase.DeleteRequest("requests", v.WalletAddress, "wallet_address")
					externaldatabase.DenyRequest(v, "invalid signature", "requests")
					continue
				}

				if !blockchain.CheckIfIsApproved(v.WalletAddress) {
					err := blockchain.ApproveAddress(v.WalletAddress)
					if err != nil {
						err = externaldatabase.DenyRequest(v, "error approving wallet on mainnet", "requests")
						if err != nil {
							embeddeddatabase.DeleteRequest("requests", v.WalletAddress, "wallet_address")
						}
					}
					err = blockchain.EnableOnPuffin(v.WalletAddress)
					if err != nil {
						err = externaldatabase.DenyRequest(v, "error enabling wallet on puffin", "requests")
						if err != nil {
							embeddeddatabase.DeleteRequest("requests", v.WalletAddress, "wallet_address")
						}
					}
					err = externaldatabase.ApproveRequest(v, "requests")
					if err == nil {
						embeddeddatabase.DeleteRequest("requests", v.WalletAddress, "wallet_address")
					}
				} else {
					err := externaldatabase.DenyRequest(v, "wallet already has kyc", "requests")
					if err != nil {
						embeddeddatabase.DeleteRequest("requests", v.WalletAddress, "wallet_address")
					}
				}
			}

			log.WithFields(log.Fields{"file": "Verify:HandleRequests", "kyc_queue_size": len(global.SubAccountQueue)}).Info("Checking Subaccount Queue")
			go global.Log(map[string]interface{}{"status": "checking queue", "message": "verifying subaccount requests", "queue_size": len(global.SubAccountQueue)})

			for _, v := range global.SubAccountQueue {
				parentIsValid := blockchain.VerifySignature(v.ParentSignature.SignatureData, v.ParentAddress)
				if !parentIsValid {
					embeddeddatabase.DeleteRequest("subaccount_requests", v.ParentAddress, "parent_wallet_address")
					externaldatabase.DenySubRequest(v, "parent invalid signature", "subaccount_requests")
					continue
				}

				subaccountIsValid := blockchain.VerifySignature(v.SubAccountSignature.SignatureData, v.SubAccountAddress)
				if !subaccountIsValid {
					embeddeddatabase.DeleteRequest("subaccount_requests", v.ParentAddress, "parent_wallet_address")
					externaldatabase.DenySubRequest(v, "subaccount invalid signature", "subaccount_requests")
					continue
				}

				if !blockchain.CheckIfIsApproved(v.SubAccountAddress) {
					err := blockchain.ApproveAddress(v.SubAccountAddress)
					if err != nil {
						err = externaldatabase.DenySubRequest(v, "error approving wallet on mainnet", "subaccount_request")
						if err != nil {
							embeddeddatabase.DeleteRequest("subaccount_requests", v.ParentAddress, "parent_wallet_address")
						}
					}
					err = blockchain.EnableOnPuffin(v.SubAccountAddress)
					if err != nil {
						err = externaldatabase.DenySubRequest(v, "error enabling wallet on puffin", "subaccount_requests")
						if err != nil {
							embeddeddatabase.DeleteRequest("subaccount_requests", v.ParentAddress, "parent_wallet_address")
						}
					}
					err = externaldatabase.ApproveSubRequest(v, "subaccount_requests")
					if err == nil {
						embeddeddatabase.DeleteRequest("subaccount_requests", v.ParentAddress, "parent_wallet_address")
					}
				} else {
					err := externaldatabase.DenySubRequest(v, "wallet already has kyc", "subaccount_requests")
					if err != nil {
						embeddeddatabase.DeleteRequest("subaccount_requests", v.ParentAddress, "parent_wallet_address")
					}
				}
			}
			updating = false
		}
	}
}

