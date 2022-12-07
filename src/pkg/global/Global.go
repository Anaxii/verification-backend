package global

var Queue []VerificationRequest
var SubAccountQueue []SubAccountRequest
var CheckRequests = make(chan bool)
var SocketChannel = make(chan interface{})
var SocketCount = 0

func Log(data interface{}) {
	if SocketCount > 0 {
		SocketChannel <- data
	}

}
