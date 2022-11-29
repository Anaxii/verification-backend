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
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	approved := externaldatabase.CheckIfExists(requestBody.WalletAddress, "approved")
	pending := externaldatabase.CheckIfExists(requestBody.WalletAddress, "requests")
	denied := externaldatabase.CheckIfExists(requestBody.WalletAddress, "denied")
	if approved {
		res, _ := json.Marshal(map[string]string{"status": "approved"})
		w.Write(res)
	} else if pending {
		res, _ := json.Marshal(map[string]string{"status": "approved"})
		w.Write(res)
	} else if denied {
		res, _ := json.Marshal(map[string]string{"status": "approved"})
		w.Write(res)
	}

}
