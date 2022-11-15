package global

import "math/big"

var Queue []VerificationRequest
var CheckRequests = make(chan bool)
var AvaxRpcURL = "https://red-weathered-firefly.avalanche-testnet.quiknode.pro/ext/bc/C/rpc"
var AvaxChainId = big.NewInt(43113)
const AvaxChainApprovedAccountsAddress = "0x094B85f01716ddB7E07bE8E68c29d1bA6E59944e"
