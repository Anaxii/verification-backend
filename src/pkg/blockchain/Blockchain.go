package blockchain

import (
	"fmt"
	"github.com/chenzhijie/go-web3"
	"github.com/chenzhijie/go-web3/types"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/common/hexutil"
	"log"
	"math/big"
	"puffinverificationbackend/src/pkg/config"
	"puffinverificationbackend/src/pkg/constants"
	"puffinverificationbackend/src/pkg/global"
	"strconv"
)

func VerifySignature(_data global.SignatureData, walletAddress string) bool  {

	web3, err := web3.NewWeb3(global.AvaxRpcURL)

	if err != nil {
		log.Println(err)
		return false
	}

	avaxMainnetChainId := int64(43113)
	if err := web3.Eth.SetAccount(config.PrivateKey); err != nil {
		log.Println(err)
		return false
	}
	web3.Eth.SetChainId(avaxMainnetChainId)
	tokenAddr := "0xF686F5D7165e8Ce1C606978F424a2DBd4a37e122"
	contract, err := web3.Eth.NewContract(constants.VerificationABI, tokenAddr)

	fmt.Println(_data.HashedMessage, _data.V, _data.R, _data.S)

	if err != nil {
		log.Println(err)
		return false
	}

	msg := [32]byte{}
	copy(msg[:], []byte(_data.HashedMessage))

	num := _data.V
	v, err := strconv.Atoi(num)
	if err != nil {
		log.Println(err)
		return false
	}

	r := [32]byte{}
	copy(r[:], []byte(_data.R))

	s := [32]byte{}
	copy(s[:], []byte(_data.S))

	signer, err := contract.Call("VerifyMessage", msg, uint8(v), r, s)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println(signer, walletAddress)

	//if signer.(common.Address).String() == walletAddress {
	//	return true
	//}
	if signer.(common.Address).String() != "0x0000000000000000000000000000000000000000" {
		return true
	}
	return false
}

func CheckIfIsApproved(walletAddress string) bool {
	web3, err := web3.NewWeb3(global.AvaxRpcURL)

	if err != nil {
		log.Println(err)
		return false
	}

	avaxMainnetChainId := int64(43113)
	if err := web3.Eth.SetAccount(config.PrivateKey); err != nil {
		log.Println(err)
		return false
	}
	web3.Eth.SetChainId(avaxMainnetChainId)
	tokenAddr := "0x094B85f01716ddB7E07bE8E68c29d1bA6E59944e"
	contract, err := web3.Eth.NewContract(constants.ApprovedAccountsABI, tokenAddr)

	isApproved, err := contract.Call("isApproved", common.HexToAddress(walletAddress))
	if err != nil {
		log.Println(err)
		return false
	}

	return isApproved.(bool)
}

func ApproveAddress(walletAddress string) bool {

	web3, err := web3.NewWeb3(global.AvaxRpcURL)

	if err != nil {
		log.Println(err)
		return false
	}

	avaxMainnetChainId := int64(43113)
	if err := web3.Eth.SetAccount(config.PrivateKey); err != nil {
		log.Println(err)
		return false
	}
	web3.Eth.SetChainId(avaxMainnetChainId)
	tokenAddr := "0x094B85f01716ddB7E07bE8E68c29d1bA6E59944e"
	contract, err := web3.Eth.NewContract(constants.ApprovedAccountsABI, "0x094B85f01716ddB7E07bE8E68c29d1bA6E59944e")
	log.Println(common.HexToAddress(walletAddress), web3.Eth.Address())

	approveInputData, err := contract.Methods("approve").Inputs.Pack(common.HexToAddress(walletAddress))
	if err != nil {
		panic(err)
	}

	tokenAddress := common.HexToAddress(tokenAddr)

	call := &types.CallMsg{
		From: web3.Eth.Address(),
		To:   tokenAddress,
		Data: approveInputData,
		Gas:  types.NewCallMsgBigInt(big.NewInt(types.MAX_GAS_LIMIT)),
	}
	// fmt.Printf("call %v\n", call)
	gasLimit, err := web3.Eth.EstimateGas(call)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Estimate gas limit %v\n", gasLimit)
	nonce, err := web3.Eth.GetNonce(web3.Eth.Address(), nil)
	if err != nil {
		panic(err)
	}
	txHash, err := web3.Eth.SyncSendRawTransaction(
		common.HexToAddress(tokenAddr),
		big.NewInt(0),
		nonce,
		gasLimit,
		web3.Utils.ToGWei(1),
		approveInputData,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Send approve tx hash %v\n", txHash)
	return false
}
