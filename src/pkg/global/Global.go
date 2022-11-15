package global

import "math/big"

var Queue []VerificationRequest
var Check = make(chan bool)
var AvaxRpcURL = "https://red-weathered-firefly.avalanche-testnet.quiknode.pro/ext/bc/C/rpc"
var AvaxChainId = big.NewInt(43113)
