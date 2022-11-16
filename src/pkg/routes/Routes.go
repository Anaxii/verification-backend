package routes

import (
	"encoding/json"
	"log"
	"net/http"
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

	id, err := externaldatabase.InsertRequest(requestBody, "requests", "pending")
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

func Status(w http.ResponseWriter, r *http.Request) {
	var requestBody global.VerificationRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	status := externaldatabase.CheckIfExists(requestBody.WalletAddress, "requests")
	res, err := json.Marshal(map[string]bool{"status": status})

	w.Write(res)

}