package global

import "math/big"

var Queue []VerificationRequest
var CheckRequests = make(chan bool)
var AvaxRpcURL = "https://red-weathered-firefly.avalanche-testnet.quiknode.pro/ext/bc/C/rpc"
var AvaxChainId = big.NewInt(43113)
const AvaxChainApprovedAccountsAddress = "0x094B85f01716ddB7E07bE8E68c29d1bA6E59944e"
const PuffinRpcURL = "https://node.thepuffin.network/ext/bc/273dwzFtrR6JQzLncTAbN5RBtiqdysVfKTJKBvYHhtUHBnrYWe/rpc"
const PuffinAllowListInterface = "0x0200000000000000000000000000000000000002"
var PuffinChainId = big.NewInt(43113114)
