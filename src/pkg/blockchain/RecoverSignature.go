// https://gist.github.com/raddy/cef9d965550b42057ce5ba7734622132
package blockchain

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

type EIP191 struct {
	msg       string
	signature string
	address   string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func hasValidLastByte(sig []byte) bool {
	return sig[64] == 0 || sig[64] == 1
}

func hasMatchingAddress(knownAddress string, recoveredAddress string) bool {
	return strings.ToLower(knownAddress) == strings.ToLower(recoveredAddress)
}

func signEIP191(message string) common.Hash {
	msg := []byte(message)
	formattedMsg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)
	return crypto.Keccak256Hash([]byte(formattedMsg))
}

func decodePersonal(eip191 EIP191) bool {
	decodedSig, err := hexutil.Decode(eip191.signature)
	check(err)

	if decodedSig[64] < 27 {
		if !hasValidLastByte(decodedSig) {
			panic("Invalid last byte")
		}
	} else {
		decodedSig[64] -= 27 // shift byte?
	}

	hash := signEIP191(eip191.msg)

	recoveredPublicKey, err := crypto.Ecrecover(hash.Bytes(), decodedSig)
	check(err)

	secp256k1RecoveredPublicKey, err := crypto.UnmarshalPubkey(recoveredPublicKey)
	check(err)

	recoveredAddress := crypto.PubkeyToAddress(*secp256k1RecoveredPublicKey).Hex()

	if hasMatchingAddress(eip191.address, recoveredAddress) {
		return true
	}
	return false
}
