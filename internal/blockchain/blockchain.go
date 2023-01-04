package blockchain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"puffinverificationbackend/internal/config"
	"puffinverificationbackend/internal/global"
	"puffinverificationbackend/pkg/abi"
)


func VerifySignature(_data global.SignatureData, walletAddress string) bool  {
	eip191 := EIP191{"Puffin KYC Request: " + walletAddress, _data.Sig, walletAddress}
	return decodePersonal(eip191)
}

func CheckIfIsApproved(walletAddress string) bool {
	conn, err := ethclient.Dial(config.AvaxRpcURL)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to connect to the Ethereum client")
	}

	verify, err := abi.NewPuffinApprovedAccounts(common.HexToAddress(config.AvaxChainApprovedAccountsAddress), conn)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to instantiate PuffinApprovedAccounts contract")
	}

	isApproved, err := verify.IsApproved(nil, common.HexToAddress(walletAddress))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to check if user is approved")
		return false
	}

	conn, err = ethclient.Dial(config.PuffinRpcURL)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to connect to puffin RPC")
		return false
	}

	verifyPuffin, err := abi.NewAllowListInterface(common.HexToAddress(config.PuffinAllowListInterfaceURL), conn)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to instantiate AllowListInterface")
		return false
	}

	isEnabled, err := verifyPuffin.ReadAllowList(nil, common.HexToAddress(walletAddress))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to call ReadAllowList")
		return false
	}

	return isApproved && isEnabled != big.NewInt(0)
}

func ApproveAddress(walletAddress string) error {
	conn, auth, err := getAuth(config.AvaxRpcURL, config.AvaxChainId)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:ApproveAddress"}).Error("Failed to get auth")
		return err
	}

	verify, err := abi.NewPuffinApprovedAccounts(common.HexToAddress(config.AvaxChainApprovedAccountsAddress), conn)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:ApproveAddress"}).Error("Failed to instantiate PuffinApprovedAccounts contract")
		return err
	}

	_, err = verify.Approve(auth, common.HexToAddress(walletAddress))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:ApproveAddress"}).Error("Failed to call approve")
		return err
	}

	return nil
}

func EnableOnPuffin(walletAddress string) error {
	conn, auth, err := getAuth(config.PuffinRpcURL, config.PuffinChainId)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:EnableOnPuffin"}).Error("Failed to get auth")
		return err
	}

	verify, err := abi.NewAllowListInterface(common.HexToAddress(config.PuffinAllowListInterfaceURL), conn)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:EnableOnPuffin"}).Error("Failed to initialize AllowListInterface")
		return err
	}

	_, err = verify.SetEnabled(auth, common.HexToAddress(walletAddress))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:EnableOnPuffin"}).Error("Failed to call SetEnabled")
		return err
	}

	return nil
}

func getAuth(rpcURL string, chainID *big.Int) (*ethclient.Client, *bind.TransactOpts, error) {

	conn, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "rpcURL": rpcURL, "chainID": chainID, "file": "Blockchain:getAuth"}).Error("Failed to connect to the Ethereum client")
		return &ethclient.Client{}, &bind.TransactOpts{}, err
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "rpcURL": rpcURL, "chainID": chainID, "file": "Blockchain:getAuth"}).Error("Failed to convert private key string to ECDSA")
		return &ethclient.Client{}, &bind.TransactOpts{}, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "rpcURL": rpcURL, "chainID": chainID, "file": "Blockchain:getAuth"}).Error("Failed to create authorized transactor")
		return &ethclient.Client{}, &bind.TransactOpts{}, err
	}

	return conn, auth, err
}

func GetTier(walletAddress string) (string, bool) {

	conn, err := ethclient.Dial(config.AvaxRpcURL)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to connect to the Ethereum client")
	}

	core, err := abi.NewPuffinCore(common.HexToAddress(config.PuffinCoreAddress), conn)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to instantiate PuffinApprovedAccounts contract")
	}

	tier, err := core.Tier(nil, common.HexToAddress(walletAddress))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to check if user is approved")
		return "0", false
	}

	isKYC, err := core.IsKYC(nil, common.HexToAddress(walletAddress))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to check if user is approved")
		return "0", false
	}
	return tier.String(), isKYC
}

func SetTier(walletAddress string, tier *big.Int) error {

	conn, auth, err := getAuth(config.AvaxRpcURL, config.AvaxChainId)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:EnableOnPuffin"}).Error("Failed to get auth")
		return err
	}

	verify, err := abi.NewPuffinCore(common.HexToAddress(config.PuffinCoreAddress), conn)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:EnableOnPuffin"}).Error("Failed to initialize AllowListInterface")
		return err
	}

	_, err = verify.SetTier(auth, common.HexToAddress(walletAddress), tier)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:EnableOnPuffin"}).Error("Failed to call SetEnabled")
		return err
	}

	return nil
}