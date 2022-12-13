package verify

import (
	"puffinverificationbackend/internal/global"
	"time"
)

func minuteTicker() *time.Ticker {
	return time.NewTicker(time.Second * time.Duration(60-time.Now().Second()))
}

func startMinuteTicker(updating *bool) {
	t := minuteTicker()
	for {
		<-t.C
		t = minuteTicker()
		if !*updating {
			global.CheckRequests <- true
		}
	}
}
