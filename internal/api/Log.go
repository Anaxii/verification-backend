package api

var SocketChannels =  make(map[string]chan interface{})

func Log(data interface{}) {
	for  k := range SocketChannels {
		SocketChannels[k] <- data
	}
}
