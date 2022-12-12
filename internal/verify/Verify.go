package verify

import (
	log "github.com/sirupsen/logrus"
	"puffinverificationbackend/internal/api"
	"puffinverificationbackend/internal/blockchain"
	"puffinverificationbackend/internal/embeddeddatabase"
	"puffinverificationbackend/internal/externaldatabase"
	"puffinverificationbackend/internal/global"
	"time"
)

func minuteTicker() *time.Ticker {
	return time.NewTicker(time.Second * time.Duration(60-time.Now().Second()))
}

func startMinuteTicker(updating *bool) {
	t := minuteTicker()
	for {
		<-t.C
		t = minuteTicker()
		if !*updating {
			global.CheckRequests <- true
		}
	}
}

func HandleRequests() {
	updating := false
	go startMinuteTicker(&updating)
	for {
		select {
		case <-global.CheckRequests:
			updating = true
			checkKYCRequests()
			updating = false
		}
	}
}

func checkKYCRequests() {
	embeddeddatabase.RefreshQueue()
	log.WithFields(log.Fields{"file": "Verify:HandleRequests", "kyc_queue_size": len(global.AccountQueue)}).Info("Checking KYC AccountQueue")
	go api.Log(map[string]interface{}{"status": "checking queue", "message": "verifying kyc requests", "queue_size": len(global.AccountQueue)})

	for _, v := range global.AccountQueue {
		handleAccountQueue(v)
	}

	log.WithFields(log.Fields{"file": "Verify:HandleRequests", "kyc_queue_size": len(global.SubAccountQueue)}).Info("Checking Subaccount AccountQueue")
	go api.Log(map[string]interface{}{"status": "checking queue", "message": "verifying subaccount requests", "queue_size": len(global.SubAccountQueue)})

	for _, v := range global.SubAccountQueue {
		handleSubaccountQueue(v)
	}
}

func handleAccountQueue(v global.AccountRequest) {
	if isValid := blockchain.VerifySignature(v.Signature.SignatureData, v.WalletAddress); !isValid {
		denyAccountAndDelete(v, "invalid signature")
		return
	}

	if blockchain.CheckIfIsApproved(v.WalletAddress) {
		denyAccountAndDelete(v, "wallet already has kyc")
		return
	}

	if err := blockchain.ApproveAddress(v.WalletAddress); err != nil {
		denyAccountAndDelete(v, "error approving wallet on mainnet")
		return
	}

	if err := blockchain.EnableOnPuffin(v.WalletAddress); err != nil {
		denyAccountAndDelete(v, "error enabling wallet on puffin")
		return
	}

	if err := externaldatabase.ApproveRequest(v, "requests"); err == nil {
		embeddeddatabase.DeleteRequest("requests", v.WalletAddress, "wallet_address")
	}

}

func handleSubaccountQueue(v global.SubAccountRequest) {
	if parentIsValid := blockchain.VerifySignature(v.ParentSignature.SignatureData, v.ParentAddress); !parentIsValid {
		denySubAccountAndDelete(v, "parent invalid signature")
		return
	}

	if subaccountIsValid := blockchain.VerifySignature(v.SubAccountSignature.SignatureData, v.SubAccountAddress); !subaccountIsValid {
		denySubAccountAndDelete(v, "subaccount invalid signature")
		return
	}

	if err := blockchain.ApproveAddress(v.SubAccountAddress); err != nil {
		denySubAccountAndDelete(v, "error approving wallet on mainnet")
		return
	}
	if err := blockchain.EnableOnPuffin(v.SubAccountAddress); err != nil {
		denySubAccountAndDelete(v, "error enabling wallet on puffin")
		return
	}

	if err := externaldatabase.ApproveSubRequest(v, "subaccount_requests"); err == nil {
		embeddeddatabase.DeleteRequest("subaccount_requests", v.ParentAddress, "parent_wallet_address")
	}
}

func denySubAccountAndDelete(v global.SubAccountRequest, reason string) {
	err := externaldatabase.DenySubRequest(v, reason, "subaccount_requests")
	if err != nil {
		embeddeddatabase.DeleteRequest("subaccount_requests", v.ParentAddress, "parent_wallet_address")
	}
}

func denyAccountAndDelete(v global.AccountRequest, reason string) {
	err := externaldatabase.DenyRequest(v, reason, "requests")
	if err != nil {
		embeddeddatabase.DeleteRequest("requests", v.WalletAddress, "wallet_address")
	}
}
