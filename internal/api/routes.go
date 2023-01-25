package api

import (
	"encoding/json"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"net/http"
	"puffinverificationbackend/internal/embeddeddatabase"
	"puffinverificationbackend/internal/externaldatabase"
	"puffinverificationbackend/internal/global"
	"puffinverificationbackend/pkg/util"
	"strings"
	"time"
)

var statusCache = cache.New(5*time.Minute, 10*time.Minute)


func verify(w http.ResponseWriter, r *http.Request) {


	var requestBody global.AccountRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		log.Println(err)
		log.WithFields(log.Fields{"error": err.Error(), "file": "Routes:verify"}).Warn("Failed to decode request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{"ip": util.ReadUserIP(r), "wallet_address": requestBody.WalletAddress}).Info("/verify")
	go Log(map[string]interface{}{"status": "kyc request", "message": "verifying if account already exists", "walletAddress": requestBody.WalletAddress})

	res, err := json.Marshal(map[string]string{"success": "true"})
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Routes:verify"}).Warn("Failed to marshal response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	approved, _ := externaldatabase.CheckIfExists(requestBody.WalletAddress, "accounts", "wallet_address")
	if approved {
		go Log(map[string]interface{}{"status": "kyc request", "message": "account already approved", "walletAddress": requestBody.WalletAddress})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	approved, _ = externaldatabase.CheckIfExists(requestBody.WalletAddress, "subaccounts", "subaccount_address")
	if approved {
		go Log(map[string]interface{}{"status": "kyc request", "message": "account already approved", "walletAddress": requestBody.WalletAddress})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestBody.Status ="pending"
	id, err := externaldatabase.InsertRequest(requestBody, "account_requests")
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Routes:verify"}).Warn("Failed to insert requestBody into external")
		go Log(map[string]interface{}{"status": "kyc request", "message": "kyc set to pending", "walletAddress": requestBody.WalletAddress})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = embeddeddatabase.InsertNewRequest(requestBody, id)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Routes:verify"}).Warn("Failed to insert requestBody into embedded")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	go Log(map[string]interface{}{"status": "kyc request", "message": "kyc set to pending", "walletAddress": requestBody.WalletAddress})

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
	go Log(map[string]interface{}{"status": "subaccount request", "message": "user requested new subaccount", "parent": requestBody.ParentAddress, "subaccount": requestBody.SubAccountAddress})

	res, err := json.Marshal(map[string]string{"success": "true"})
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Routes:verify"}).Warn("Failed to marshal response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	approved, _ := externaldatabase.CheckIfExists(requestBody.ParentAddress, "accounts", "wallet_address")
	if !approved {
		res, _ := json.Marshal(map[string]string{"status": "parentAddressNotExist"})
		go Log(map[string]interface{}{"status": "subaccount request", "message": "parent address invalid", "parent": requestBody.ParentAddress, "subaccount": requestBody.SubAccountAddress})
		w.Write(res)
		return
	}

	approved, _ = externaldatabase.CheckIfExists(requestBody.SubAccountAddress, "subaccounts", "subaccount_address")
	if approved {
		res, _ := json.Marshal(map[string]string{"status": "subaccountAlreadyClaimed"})
		go Log(map[string]interface{}{"status": "subaccount request", "message": "subaccount already claimed", "parent": requestBody.ParentAddress, "subaccount": requestBody.SubAccountAddress})
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
	go Log(map[string]interface{}{"status": "subaccount request", "message": "subaccount request set to pending", "parent": requestBody.ParentAddress, "subaccount": requestBody.SubAccountAddress})

	global.CheckRequests <- true

	w.Write(res)

}

func status(w http.ResponseWriter, r *http.Request) {
	var requestBody global.AccountRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Routes:status"}).Warn("Failed to decode request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	val, found := statusCache.Get(requestBody.WalletAddress)
	if found {
		log.Println(val)
		var res []byte
		if !val.(bool) {
			res, _ = json.Marshal(map[string]string{"status": "sub"})
		} else {
			res, _ = json.Marshal(map[string]string{"status": "approved"})
		}
		w.Write(res)
		return
	}

	if requestBody.WalletAddress == "" {
		log.WithFields(log.Fields{"file": "Routes:status"}).Warn("User didnt provide address")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestBody.WalletAddress = strings.ToLower(requestBody.WalletAddress)

	log.WithFields(log.Fields{"ip": util.ReadUserIP(r), "wallet_address": requestBody.WalletAddress}).Info("/status")

	exists, data := externaldatabase.CheckIfExists(requestBody.WalletAddress, "account_requests", "wallet_address")
	if !exists {
		res, _ := json.Marshal(map[string]string{"status": "nonExist"})
		w.Write(res)
		return
	}
	if data.Status == "approved" {
		res, _ := json.Marshal(map[string]string{"status": "approved"})
		w.Write(res)
		go Log(map[string]interface{}{"status": "kyc status request", "message": "account verified", "walletAddress": requestBody.WalletAddress})
		go statusCache.Set(requestBody.WalletAddress, true, cache.DefaultExpiration)
		return
	}
	if data.Status == "pending" {
		res, _ := json.Marshal(map[string]string{"status": "approved"})
		w.Write(res)
		go Log(map[string]interface{}{"status": "kyc status request", "message": "account pending", "walletAddress": requestBody.WalletAddress})
		return
	}
	if data.Status == "denied" {
		res, _ := json.Marshal(map[string]string{"status": "denied"})
		w.Write(res)
		go Log(map[string]interface{}{"status": "kyc status request", "message": "account denied", "walletAddress": requestBody.WalletAddress})
		return
	}

	exists, _ = externaldatabase.CheckIfExists(requestBody.WalletAddress, "subaccounts", "wallet_address")
	if exists {
		res, _ := json.Marshal(map[string]string{"status": "sub"})
		w.Write(res)
		go Log(map[string]interface{}{"status": "kyc status request", "message": "account is sub", "walletAddress": requestBody.WalletAddress})
		return
	}

	res, _ := json.Marshal(map[string]string{"status": "nonExist"})
	w.Write(res)
}

func geoTier(w http.ResponseWriter, r *http.Request) {
	var requestBody global.AccountRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Routes:status"}).Warn("Failed to decode request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if requestBody.WalletAddress == "" {
		log.WithFields(log.Fields{"file": "Routes:status"}).Warn("User didnt provide address")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	approved, data := externaldatabase.CheckIfExists(requestBody.WalletAddress, "approved", "wallet_address")
	blocked := []string{"united states", "usa"}
	if approved {
		for _, a := range blocked {
			if strings.ToLower(data.Country) == a {
				res, _ := json.Marshal(map[string]int{"tier": 1})
				w.Write(res)
				return
			}
		}
		res, _ := json.Marshal(map[string]int{"tier": 0})
		w.Write(res)
		return
	}

	res, _ := json.Marshal(map[string]string{"error": "nonExist"})
	w.Write(res)
	return

}

func getWS(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	data := make(chan interface{})
	id := util.RandStringRunes(20)
	SocketChannels[id] = data
	reader(ws, data, id)
}