package main

import (
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"puffinverificationbackend/src/pkg/Log"
	"puffinverificationbackend/src/pkg/api"
	"puffinverificationbackend/src/pkg/embeddeddatabase"
	"puffinverificationbackend/src/pkg/global"
	"puffinverificationbackend/src/pkg/verify"
)

func main() {

	Log.SetupLogs()

	log.Info("Logging configured")
	embeddeddatabase.SetupDatabase()
	log.Info("Database Setup")
	log.Info("Starting Request Handler")
	go verify.HandleRequests()
	global.CheckRequests <- true

	log.Info("Starting API")
	api.StartAPI()

}
