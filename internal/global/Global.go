package global

var AccountQueue []AccountRequest
var SubAccountQueue []SubAccountRequest
var CheckRequests = make(chan bool)
