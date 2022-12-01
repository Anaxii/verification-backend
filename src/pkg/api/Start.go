package api

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"puffinverificationbackend/src/pkg/config"
)

func StartAPI() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST"},
	})

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/verify", verify).Methods("POST")
	r.HandleFunc("/requestsubaccount", requestSubaccount).Methods("POST")
	r.HandleFunc("/status", status).Methods("POST")
	r.HandleFunc("/pub", getPub).Methods("GET")

	r.Use(mux.CORSMethodMiddleware(r))
	log.Fatal(http.ListenAndServe(":"+config.Port, c.Handler(r)))
}
