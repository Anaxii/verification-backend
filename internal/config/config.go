package config

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
	"io/ioutil"
	"math/big"
	"os"
	"puffinverificationbackend/internal/global"
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
var PuffinCoreAddress string

func init() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.WithFields(log.Fields{"error": err.Error(), "file": "config:init"}).Fatal("Failed to generate a private key")
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
			PuffinCoreAddress: "0x44b4cCbdf70325f0a6cf8644ecf2AdBf7c737329",
			PuffinChainID:                    43113114,
		}, "", "  ")
		_ = ioutil.WriteFile("config.json", file, 0644)
		log.WithFields(log.Fields{ "file": "config:init"}).Fatal("Generated config.json | Fill in empty data and run again")

	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "config:init"}).Fatal("Config file is invalid")
	}

	var config global.ConfigStruct
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "config:init"}).Fatal("Could not unmarshal config file")
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
	PuffinCoreAddress = config.PuffinCoreAddress
}

func GenerateECDSAKey(pkey string) (string, *ecdsa.PrivateKey) {
	privateKey, err := crypto.HexToECDSA(pkey)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "config:init"}).Error("Could not convert private key string to ECDSA")
		return "", &ecdsa.PrivateKey{}
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
