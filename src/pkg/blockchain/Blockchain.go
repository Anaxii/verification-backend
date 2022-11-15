package blockchain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"puffinverificationbackend/src/pkg/abi"
	"puffinverificationbackend/src/pkg/config"
	"puffinverificationbackend/src/pkg/global"
)


func VerifySignature(_data global.SignatureData, walletAddress string) bool  {
	eip191 := EIP191{"Puffin KYC Request: " + walletAddress, _data.Sig, walletAddress}
	return decodePersonal(eip191)
}

func CheckIfIsApproved(walletAddress string) bool {
	conn, err := ethclient.Dial(global.AvaxRpcURL)
	if err != nil {
		log.Println("Failed to connect to the Ethereum client:", err)
	}

	verify, err := abi.NewPuffinApprovedAccounts(common.HexToAddress(global.AvaxChainApprovedAccountsAddress), conn)
	if err != nil {
		log.Println("Failed to instantiate PuffinApprovedAccounts contract:", err)
	}

	isApproved, err := verify.IsApproved(nil, common.HexToAddress(walletAddress))
	if err != nil {
		log.Println("Failed to read user:", err)
		return false
	}

	conn, err = ethclient.Dial(global.PuffinRpcURL)
	if err != nil {
		log.Println("Failed to connect to the Ethereum client:", err)
	}

	verifyPuffin, err := abi.NewAllowListInterface(common.HexToAddress(global.PuffinAllowListInterface), conn)
	if err != nil {
		log.Println("Failed to instantiate PuffinApprovedAccounts contract:", err)
	}

	isEnabled, err := verifyPuffin.ReadAllowList(nil, common.HexToAddress(walletAddress))
	if err != nil {
		log.Println("Failed to read user:", err)
		return false
	}

	return isApproved && isEnabled != big.NewInt(0)
}

func ApproveAddress(walletAddress string) error {

	conn, err := ethclient.Dial(global.AvaxRpcURL)
	if err != nil {
		log.Println("Failed to connect to the Ethereum client:", err)
		return err
	}

	verify, err := abi.NewPuffinApprovedAccounts(common.HexToAddress(global.AvaxChainApprovedAccountsAddress), conn)
	if err != nil {
		log.Println("Failed to instantiate PuffinApprovedAccounts contract:", err)
		return err
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Println(err)
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, global.AvaxChainId)
	if err != nil {
		log.Println("Failed to create authorized transactor:", err)
		return err
	}

	_, err = verify.Approve(auth, common.HexToAddress(walletAddress))
	if err != nil {
		log.Println("Failed to update user:", err)
		return err
	}

	return nil
}

func EnableOnPuffin(walletAddress string) error {

	conn, err := ethclient.Dial(global.PuffinRpcURL)
	if err != nil {
		log.Println("Failed to connect to the Ethereum client:", err)
		return err
	}

	verify, err := abi.NewAllowListInterface(common.HexToAddress(global.PuffinAllowListInterface), conn)
	if err != nil {
		log.Println("Failed to instantiate PuffinApprovedAccounts contract:", err)
		return err
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Println(err)
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, global.PuffinChainId)
	if err != nil {
		log.Println("Failed to create authorized transactor:", err)
		return err
	}

	_, err = verify.SetEnabled(auth, common.HexToAddress(walletAddress))
	if err != nil {
		log.Println("Failed to update user:", err)
		return err
	}

	return nil
}
