package global

type VerificationRequest struct {
	Name            string          `json:"name" bson:"name"`
	ID              string          `json:"id" bson:"id"`
	Status          string          `json:"status" bson:"status"`
	Email           string          `json:"email" bson:"email"`
	WalletAddress   string          `json:"wallet_address" bson:"wallet_address"`
	Country         string          `json:"country" bson:"country"`
	PhysicalAddress string          `json:"physical_address" bson:"physical_address"`
	IdentityNumber  string          `json:"identity_number" bson:"identity_number"`
	DateOfBirth     string          `json:"date_of_birth" bson:"date_of_birth"`
	Signature       SignatureStruct `json:"signature" bson:"signature"`
	Beneficiary     struct {
		Name          string `json:"name" bson:"name"`
		WalletAddress string `json:"wallet_address" bson:"wallet_address"`
	}
}

type SignatureStruct struct {
	Message       string        `json:"message" bson:"message"`
	Account       string        `json:"account" bson:"account"`
	SignatureData SignatureData `json:"signature_data" bson:"signature_data"`
}

type SignatureData struct {
	HashedMessage string `json:"hashed_message" bson:"hashed_message"`
	R             string `json:"r" bson:"r"`
	S             string `json:"s" bson:"s"`
	V             string `json:"v" bson:"v"`
	Sig           string `json:"sig" bson:"sig"`
}

type ConfigStruct struct {
	PrivateKey                       string `json:"private_key"`
	Port                             string `json:"port"`
	MongoDbURI                       string `json:"mongo_db_uri"`
	AvaxRPCURL                       string `json:"avax_rpc_url"`
	AvaxChainID                      int64  `json:"avax_chain_id"`
	AvaxChainApprovedAccountsAddress string `json:"avax_chain_approved_accounts_address"`
	PuffinRPCURL                     string `json:"puffin_rpc_url"`
	PuffinAllowListInterfaceURL      string `json:"puffin_allow_list_interface_url"`
	PuffinChainID                    int64  `json:"puffin_chain_id"`
}

type SubAccountRequest struct {
	ParentAddress       string          `json:"parent_address" bson:"parent_address"`
	SubAccountAddress   string          `json:"subaccount_address" bson:"subaccount_address"`
	ID                  string          `json:"id" bson:"id"`
	Status              string          `json:"status" bson:"status"`
	ParentSignature     SignatureStruct `json:"parent_signature" bson:"parent_signature"`
	SubAccountSignature SignatureStruct `json:"sub_account_signature" bson:"sub_account_signature"`
}
