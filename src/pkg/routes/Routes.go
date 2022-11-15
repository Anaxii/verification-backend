package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"puffinverificationbackend/src/pkg/database"
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

	err = database.InsertNewRequest(requestBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	global.Check <- true

	w.Write(res)
}