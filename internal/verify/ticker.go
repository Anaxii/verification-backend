package verify

import (
	"puffinverificationbackend/internal/global"
	"time"
)

func minuteTicker(seconds int) *time.Ticker {
	return time.NewTicker(time.Second * time.Duration(seconds-time.Now().Second()))
}

func startMinuteTicker(updating *bool) {
	requestsTicker := minuteTicker(60)
	//countryTicker := minuteTicker(300)
	go func() {
		for {
			<-requestsTicker.C
			requestsTicker = minuteTicker(60)
			if !*updating {
				global.CheckRequests <- true
			}
		}
	}()
	//go func() {
	//for {
	//	<-countryTicker.C
	//	countryTicker = minuteTicker(300)
	//	countries.CheckCountries()
	//}
	//}()
}
