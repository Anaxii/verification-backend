package config

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"puffinverificationbackend/src/pkg/global"
)

var PrivateKey string
var PublicKey string
var Port string
var MongoDBURI string
var AvaxRpcURL string
var AvaxChainId *big.Int
var AvaxChainApprovedAccountsAddress string
var PuffinRpcURL string
var PuffinAllowListInterfaceURL string
var PuffinChainId *big.Int

func init() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}
		privateKeyBytes := crypto.FromECDSA(privateKey)

		file, _ := json.MarshalIndent(global.ConfigStruct{
			PrivateKey:                       fmt.Sprintf("%v", hexutil.Encode(privateKeyBytes)[2:]),
			Port:                             "80",
			AvaxRPCURL:                       "https://red-weathered-firefly.avalanche-testnet.quiknode.pro/ext/bc/C/rpc",
			AvaxChainID:                      43113,
			AvaxChainApprovedAccountsAddress: "0x094B85f01716ddB7E07bE8E68c29d1bA6E59944e",
			PuffinRPCURL:                     "https://node.thepuffin.network/ext/bc/273dwzFtrR6JQzLncTAbN5RBtiqdysVfKTJKBvYHhtUHBnrYWe/rpc",
			PuffinAllowListInterfaceURL:      "0x0200000000000000000000000000000000000002",
			PuffinChainID:                    43113114,
		}, "", "  ")
		_ = ioutil.WriteFile("config.json", file, 0644)
		log.Fatal("Generated config.json | Fill in empty data and run again")
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Config JSON invalid", err)
	}

	var config global.ConfigStruct
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatal("Could not parse config", err)
	}

	PrivateKey = config.PrivateKey

	if config.PrivateKey != "" {
		_publicKey, _ := GenerateECDSAKey(config.PrivateKey)
		PublicKey = _publicKey
	}

	Port = config.Port
	MongoDBURI = config.MongoDbURI
	AvaxRpcURL = config.AvaxRPCURL
	AvaxChainId = big.NewInt(config.AvaxChainID)
	AvaxChainApprovedAccountsAddress = config.AvaxChainApprovedAccountsAddress
	PuffinRpcURL = config.PuffinRPCURL
	PuffinAllowListInterfaceURL = config.PuffinAllowListInterfaceURL
	PuffinChainId = big.NewInt(config.PuffinChainID)
}

func GenerateECDSAKey(pkey string) (string, *ecdsa.PrivateKey) {
	privateKey, err := crypto.HexToECDSA(pkey)
	if err != nil {
		log.Println(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	_publicKey := hexutil.Encode(hash.Sum(nil)[12:])
	_privateKey := privateKey

	return _publicKey, _privateKey
}
