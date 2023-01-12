package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"puffinverificationbackend/internal/config"
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

	r.HandleFunc("/kyc/verify", verify).Methods("POST")
	r.HandleFunc("/kyc/setcountry", setCountry).Methods("POST")
	r.HandleFunc("/kyc/geotier", geoTier).Methods("POST")
	r.HandleFunc("/kyc/requestsubaccount", requestSubaccount).Methods("POST")
	r.HandleFunc("/kyc/status", status).Methods("POST")
	r.HandleFunc("/ws", getWS).Methods("GET")

	r.Use(mux.CORSMethodMiddleware(r))
	log.Info(fmt.Sprintf("API listening on port %v", config.Port))
	log.Fatal(http.ListenAndServe(":"+config.Port, c.Handler(r)))
}
