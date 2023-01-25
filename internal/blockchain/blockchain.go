package blockchain

import (
	_ "github.com/ethereum/go-ethereum/common/hexutil"
	"puffinverificationbackend/internal/global"
)


func VerifySignature(_data global.SignatureData, walletAddress string, message string) bool  {
	eip191 := EIP191{message, _data.Sig, walletAddress}
	return decodePersonal(eip191)
}