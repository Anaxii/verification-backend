package api

import (
	"encoding/json"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"net/http"
	"puffinverificationbackend/src/pkg/config"
	"puffinverificationbackend/src/pkg/embeddeddatabase"
	"puffinverificationbackend/src/pkg/externaldatabase"
	"puffinverificationbackend/src/pkg/global"
	"puffinverificationbackend/src/pkg/util"
	"time"
)

var statusCache = cache.New(5*time.Minute, 10*time.Minute)


func verify(w http.ResponseWriter, r *http.Request) {


	var requestBody global.VerificationRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Routes:verify"}).Warn("Failed to decode request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{"ip": util.ReadUserIP(r), "wallet_address": requestBody.WalletAddress}).Info("/verify")

	res, err := json.Marshal(map[string]string{"success": "true"})
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Routes:verify"}).Warn("Failed to marshal response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	approved, _ := externaldatabase.CheckIfExists(requestBody.WalletAddress, "approved", "wallet_address")
	if !approved {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	approved, _ = externaldatabase.CheckIfExists(requestBody.WalletAddress, "subaccounts", "subaccount_address")
	if approved {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestBody.Status ="pending"
	id, err := externaldatabase.InsertRequest(requestBody, "requests")
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Routes:verify"}).Warn("Failed to insert requestBody into external")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = embeddeddatabase.InsertNewRequest(requestBody, id)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Routes:verify"}).Warn("Failed to insert requestBody into embedded")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	global.CheckRequests <- true

	w.Write(res)

}

func requestSubaccount(w http.ResponseWriter, r *http.Request) {
	var requestBody global.SubAccountRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Routes:requestSubaccount"}).Warn("Failed to decode request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{"ip": util.ReadUserIP(r), "parent_address": requestBody.ParentAddress, "subaccount_address": requestBody.SubAccountAddress}).Info("/requestsubaccount")

	res, err := json.Marshal(map[string]string{"success": "true"})
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Routes:verify"}).Warn("Failed to marshal response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	approved, _ := externaldatabase.CheckIfExists(requestBody.ParentAddress, "approved", "wallet_address")
	if !approved {
		res, _ := json.Marshal(map[string]string{"status": "parentAddressNotExist"})
		w.Write(res)
		return
	}

	approved, _ = externaldatabase.CheckIfExists(requestBody.SubAccountAddress, "approved", "wallet_address")
	if approved {
		res, _ := json.Marshal(map[string]string{"status": "subaccountAlreadyKYC"})
		w.Write(res)
		return
	}

	approved, _ = externaldatabase.CheckIfExists(requestBody.SubAccountAddress, "subaccounts", "subaccount_address")
	if approved {
		res, _ := json.Marshal(map[string]string{"status": "subaccountAlreadyClaimed"})
		w.Write(res)
		return
	}

	id, err := externaldatabase.InsertRequest(requestBody, "subaccount_requests")
	if err != nil {
		log.Println(err)
		log.WithFields(log.Fields{"error": err.Error(), "file": "Routes:requestSubaccount"}).Warn("Failed to insert requestBody to external")
		return
	}

	err = embeddeddatabase.InsertNewSubAccountRequest(requestBody, id)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Routes:requestSubaccount"}).Warn("Failed to insert requestBody to embedded")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	global.CheckRequests <- true

	w.Write(res)

}

func getPub(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	log.WithFields(log.Fields{"ip": util.ReadUserIP(r)}).Info("/pub")

	data, err := json.Marshal(map[string]string{"pub": config.PublicKey})
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Routes:getPub"}).Warn("Failed to insert requestBody to embedded")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(data)

}

func status(w http.ResponseWriter, r *http.Request) {
	var requestBody global.VerificationRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Routes:status"}).Warn("Failed to decode request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, found := statusCache.Get(requestBody.WalletAddress)
	if found {
		res, _ := json.Marshal(map[string]string{"status": "approved"})
		w.Write(res)
		return
	}

	if requestBody.WalletAddress == "" {
		log.WithFields(log.Fields{"file": "Routes:status"}).Warn("User didnt provide address")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.WithFields(log.Fields{"ip": util.ReadUserIP(r), "wallet_address": requestBody.WalletAddress}).Info("/status")

	approved, _ := externaldatabase.CheckIfExists(requestBody.WalletAddress, "approved", "wallet_address")
	if approved {
		res, _ := json.Marshal(map[string]string{"status": "approved"})
		w.Write(res)
		statusCache.Set(requestBody.WalletAddress, true, cache.DefaultExpiration)
		return
	}
	approved, _ = externaldatabase.CheckIfExists(requestBody.WalletAddress, "subaccounts", "subaccount_address")

	if approved {
		res, _ := json.Marshal(map[string]string{"status": "approved"})
		w.Write(res)
		statusCache.Set(requestBody.WalletAddress, true, cache.DefaultExpiration)
		return
	}
	pending, _ := externaldatabase.CheckIfExists(requestBody.WalletAddress, "requests", "wallet_address")
	if pending {
		res, _ := json.Marshal(map[string]string{"status": "approved"})
		w.Write(res)
		return
	}
	denied, _ := externaldatabase.CheckIfExists(requestBody.WalletAddress, "denied", "wallet_address")
	if denied {
		res, _ := json.Marshal(map[string]string{"status": "approved"})
		w.Write(res)
		return
	}
	res, _ := json.Marshal(map[string]string{"status": "nonExist"})
	w.Write(res)
	return

}
