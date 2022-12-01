package main

import (
	_ "github.com/mattn/go-sqlite3"
	"puffinverificationbackend/src/pkg/api"
	"puffinverificationbackend/src/pkg/embeddeddatabase"
	"puffinverificationbackend/src/pkg/global"
	"puffinverificationbackend/src/pkg/verify"
)

func main() {

	embeddeddatabase.SetupDatabase()
	go verify.HandleRequests()
	global.CheckRequests <- true

	api.StartAPI()
}
