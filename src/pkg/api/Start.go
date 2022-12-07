package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"puffinverificationbackend/src/pkg/config"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 30,
	WriteBufferSize: 1024 * 30,
}

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
	r.HandleFunc("/ws", getWS).Methods("GET")

	log.WithFields(log.Fields{"verify": "POST", "requestsubaccount": "POST", "status": "POST", "pub": "GET"}).Info("Available endpoints")
	r.Use(mux.CORSMethodMiddleware(r))
	log.Info(fmt.Sprintf("API listening on port %v", config.Port))
	log.Fatal(http.ListenAndServe(":"+config.Port, c.Handler(r)))
}
