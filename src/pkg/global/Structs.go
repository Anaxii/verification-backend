package global

type VerificationRequest struct {
	Name            string          `json:"name"`
	Email           string          `json:"email"`
	WalletAddress   string          `json:"wallet_address"`
	PhysicalAddress string          `json:"physical_address"`
	IdentityNumber  string          `json:"identity_number"`
	DateOfBirth     string          `json:"date_of_birth"`
	Signature       SignatureStruct `json:"signature"`
	Beneficiary     struct {
		Name          string `json:"name"`
		WalletAddress string `json:"wallet_address"`
	}
}

type SignatureStruct struct {
	Message       string        `json:"message"`
	Account       string        `json:"account"`
	SignatureData SignatureData `json:"signature_data"`
}

type SignatureData struct {
	HashedMessage string `json:"hashed_message"`
	R             string `json:"r"`
	S             string `json:"s"`
	V             string `json:"v"`
	Sig           string `json:"sig"`
}
