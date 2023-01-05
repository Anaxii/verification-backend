package global

type CountryRequest struct {
	Country string `json:"country" bson:"country"`
	Allowed bool   `json:"allowed" bson:"allowed"`
}

type Countries struct {
	Countries []string `json:"countries" bson:"countries"`
}

type AccountRequest struct {
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
	PuffinCoreAddress                string `json:"puffin_core_address"`
}

type SubAccountRequest struct {
	ParentAddress       string          `json:"parent_address" bson:"parent_address"`
	SubAccountAddress   string          `json:"subaccount_address" bson:"subaccount_address"`
	ID                  string          `json:"id" bson:"id"`
	Status              string          `json:"status" bson:"status"`
	ParentSignature     SignatureStruct `json:"parent_signature" bson:"parent_signature"`
	SubAccountSignature SignatureStruct `json:"sub_account_signature" bson:"sub_account_signature"`
}


type ClientSettings struct {
	AdminName                   string   `json:"admin_name" bson:"admin_name"`
	ProjectName                 string   `json:"project_name" bson:"project_name"`
	AdminTelegram               string   `json:"admin_telegram" bson:"admin_telegram"`
	TeamSize                    int      `json:"team_size" bson:"team_size"`
	ProjectCommunicationChannel string   `json:"project_communication_channel" bson:"project_communication_channel"`
	ProjectWebsite              string   `json:"project_website" bson:"project_website"`
	DappName                    string   `json:"dapp_name" bson:"dapp_name"`
	RPCURL                      string   `json:"rpc_url" bson:"rpc_url"`
	GasTokenSymbol              string   `json:"gas_token_symbol" bson:"gas_token_symbol"`
	ChainID                     int      `json:"chain_id" bson:"chain_id"`
	VM                          string   `json:"vm" bson:"vm"`
	Package                     string   `json:"package" bson:"package"`
	PackageOptions              []string `json:"package_options" bson:"package_options"`
	MaxUsers                    int      `json:"max_users" bson:"max_users"`
	PuffinClientAddress         string   `json:"puffin_client_address" bson:"puffin_client_address"`
	PuffinCoreAddress           string   `json:"puffin_core_address" bson:"puffin_core_address"`
	AdminWalletAddress          string   `json:"admin_wallet_address" bson:"admin_wallet_address"`
	UUID                        int      `json:"UUID" bson:"UUID"`
	Status                      string   `json:"status" bson:"status"`
	PaymentExpiration           int      `json:"payment_expiration" bson:"payment_expiration"`
}