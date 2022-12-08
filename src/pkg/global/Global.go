package global

var Queue []VerificationRequest
var SubAccountQueue []SubAccountRequest
var CheckRequests = make(chan bool)

var SocketChannels =  make(map[string]chan interface{})

func Log(data interface{}) {
	for  k := range SocketChannels {
		SocketChannels[k] <- data
	}
}
