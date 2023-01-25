package kyc

import (
	chainanalysis "github.com/soloth/go-chainanalysis/client"
	"puffinverificationbackend/internal/global"
)

func CheckKYC(v global.AccountRequest) (string, error) {
	isSanctioned, _, err := chainanalysis.NewClient().UseDefault().IsSanctionedConcurrent(v.WalletAddress)
	if err != nil  {
		return "wait", err
	}
	if !isSanctioned {
		return "denied", err
	}
	return "approved", err
}
