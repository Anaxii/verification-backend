package verify

import (
	"puffinverificationbackend/internal/blockchain"
	"puffinverificationbackend/internal/embeddeddatabase"
	"puffinverificationbackend/internal/externaldatabase"
	"puffinverificationbackend/internal/global"
)

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
