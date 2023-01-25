package verify

import (
	"puffinverificationbackend/internal/blockchain"
	"puffinverificationbackend/internal/embeddeddatabase"
	"puffinverificationbackend/internal/externaldatabase"
	"puffinverificationbackend/internal/global"
	"puffinverificationbackend/internal/kyc"
)

func handleAccountQueue(v global.AccountRequest) {
	if isValid := blockchain.VerifySignature(v.Signature.SignatureData, v.WalletAddress, v.Signature.Message); !isValid {
		denyAccountAndDelete(v, "invalid signature")
		return
	}

	status, _ := kyc.CheckKYC(v)
	if status == "approved" {
		if err := externaldatabase.ApproveRequest(v, "account_requests"); err == nil {
			embeddeddatabase.DeleteRequest("requests", v.WalletAddress, "wallet_address")
		}
	} else if status == "wait" {
		return
	} else if status == "denied" {
		denyAccountAndDelete(v, "kyc denied")
		return
	}

}

func handleSubaccountQueue(v global.SubAccountRequest) {
	if parentIsValid := blockchain.VerifySignature(v.ParentSignature.SignatureData, v.ParentAddress, v.ParentSignature.Message); !parentIsValid {
		denySubAccountAndDelete(v, "parent invalid signature")
		return
	}

	if subaccountIsValid := blockchain.VerifySignature(v.SubAccountSignature.SignatureData, v.SubAccountAddress, v.SubAccountSignature.Message); !subaccountIsValid {
		denySubAccountAndDelete(v, "subaccount invalid signature")
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
	err := externaldatabase.DenyRequest(v, reason, "account_requests")
	if err != nil {
		embeddeddatabase.DeleteRequest("account_requests", v.WalletAddress, "wallet_address")
	}
}
