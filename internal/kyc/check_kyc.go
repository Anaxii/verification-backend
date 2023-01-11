package kyc

import "puffinverificationbackend/internal/global"

func CheckKYC(v global.AccountRequest) string {
	// handle kyc checks
	return "approved"
}
