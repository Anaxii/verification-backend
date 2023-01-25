package main

import (
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"puffinverificationbackend/internal/api"
	"puffinverificationbackend/internal/embeddeddatabase"
	"puffinverificationbackend/internal/global"
	"puffinverificationbackend/internal/verify"
	puffinLog "puffinverificationbackend/pkg/log"
)

func main() {

	puffinLog.SetupLogs()

	log.Info("Logging configured")

	embeddeddatabase.SetupDatabase()
	log.Info("Database Setup")

	log.Info("Starting Request Handler")
	go verify.HandleRequests()
	global.CheckRequests <- true

	log.Info("Starting API")
	api.StartAPI()

}
