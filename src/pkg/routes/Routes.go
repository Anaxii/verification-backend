package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"puffinverificationbackend/src/pkg/config"
	"puffinverificationbackend/src/pkg/embeddeddatabase"
	"puffinverificationbackend/src/pkg/externaldatabase"
	"puffinverificationbackend/src/pkg/global"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	var requestBody global.VerificationRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(map[string]string{"success": "true"})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	requestBody.Status ="pending"
	id, err := externaldatabase.InsertRequest(requestBody, "requests")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = embeddeddatabase.InsertNewRequest(requestBody, id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	global.CheckRequests <- true

	w.Write(res)

}

func RequestSubaccount(w http.ResponseWriter, r *http.Request) {
	var requestBody global.SubAccountRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(map[string]string{"success": "true"})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	approved, _ := externaldatabase.CheckIfExists(requestBody.ParentAddress, "approved")
	if !approved {
		res, _ := json.Marshal(map[string]string{"status": "parentAddressNotExist"})
		w.Write(res)
		return
	}

	approved, _ = externaldatabase.CheckIfExists(requestBody.SubAccountAddress, "approved")
	if approved {
		res, _ := json.Marshal(map[string]string{"status": "alreadyExist"})
		w.Write(res)
		return
	}

	id, err := externaldatabase.InsertRequest(requestBody, "subaccountRequests")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = embeddeddatabase.InsertNewSubAccountRequest(requestBody, id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	global.CheckRequests <- true

	w.Write(res)

}

func GetPub(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	data, err := json.Marshal(map[string]string{"pub": config.PublicKey})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(data)

}

func Status(w http.ResponseWriter, r *http.Request) {
	var requestBody global.VerificationRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		log.Println("cant decode", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if requestBody.WalletAddress == "" {
		log.Println("no addy", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	approved, _ := externaldatabase.CheckIfExists(requestBody.WalletAddress, "approved")
	if approved {
		res, _ := json.Marshal(map[string]string{"status": "approved"})
		w.Write(res)
		return
	}
	pending, _ := externaldatabase.CheckIfExists(requestBody.WalletAddress, "requests")
	if pending {
		res, _ := json.Marshal(map[string]string{"status": "approved"})
		w.Write(res)
		return
	}
	denied, _ := externaldatabase.CheckIfExists(requestBody.WalletAddress, "denied")
	if denied {
		res, _ := json.Marshal(map[string]string{"status": "approved"})
		w.Write(res)
		return
	}
	res, _ := json.Marshal(map[string]string{"status": "nonExist"})
	w.Write(res)
	return

}
