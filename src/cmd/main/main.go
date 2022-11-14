package main

import (
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
	"log"
	"net/http"
	"puffinverificationbackend/src/pkg/config"
	"puffinverificationbackend/src/pkg/database"
	"puffinverificationbackend/src/pkg/global"
	"puffinverificationbackend/src/pkg/routes"
	"puffinverificationbackend/src/pkg/verify"
)

func main() {

	database.SetupDatabase()

	go verify.HandleRequests()
	global.Check <- true

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},    // All origins
		AllowedMethods: []string{"POST"}, // Allowing only get, just an example
	})

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/verify", routes.Verify).Methods("POST")
	r.Use(mux.CORSMethodMiddleware(r))
	log.Fatal(http.ListenAndServe(":" + config.Port, c.Handler(r)))
}









