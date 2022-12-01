package global

var Queue []VerificationRequest
var SubAccountQueue []SubAccountRequest
var CheckRequests = make(chan bool)
