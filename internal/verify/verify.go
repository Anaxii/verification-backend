package verify

import (
	log "github.com/sirupsen/logrus"
	"puffinverificationbackend/internal/api"
	"puffinverificationbackend/internal/countries"
	"puffinverificationbackend/internal/embeddeddatabase"
	"puffinverificationbackend/internal/global"
)

func HandleRequests() {
	updating := false
	go startMinuteTicker(&updating)
	countries.CheckCountries()
	for {
		select {
		case <-global.CheckRequests:
			updating = true
			checkKYCRequests()
			updating = false
		}
	}
}

func checkKYCRequests() {
	embeddeddatabase.RefreshQueue()
	log.WithFields(log.Fields{"file": "Verify:HandleRequests", "kyc_queue_size": len(global.AccountQueue)}).Info("Checking KYC AccountQueue")
	go api.Log(map[string]interface{}{"status": "checking queue", "message": "verifying kyc requests", "queue_size": len(global.AccountQueue)})

	for _, v := range global.AccountQueue {
		handleAccountQueue(v)
	}

	log.WithFields(log.Fields{"file": "Verify:HandleRequests", "kyc_queue_size": len(global.SubAccountQueue)}).Info("Checking Subaccount AccountQueue")
	go api.Log(map[string]interface{}{"status": "checking queue", "message": "verifying subaccount requests", "queue_size": len(global.SubAccountQueue)})

	for _, v := range global.SubAccountQueue {
		handleSubaccountQueue(v)
	}
}
