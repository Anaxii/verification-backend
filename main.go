package main

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
	"log"
	"math/big"
	"net/http"
	"os"
)

const VerificationABI string = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hashedMessage\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"VerifyMessage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"
const PrivateKey string = "e5fb0910de1ba57e0328d591343c47e2ed620bdc7ba9364b62b1ff32968c45f7"


type VerificationRequest struct {
	Name            string          `json:"name"`
	Email           string          `json:"email"`
	WalletAddress   string          `json:"wallet_address"`
	PhysicalAddress string          `json:"physical_address"`
	IdentityNumber  string          `json:"identity_number"`
	DateOfBirth     string          `json:"date_of_birth"`
	Signature       SignatureStruct `json:"signature"`
	Beneficiary     struct {
		Name          string `json:"name"`
		WalletAddress string `json:"wallet_address"`
	}
}

type SignatureStruct struct {
	Message       string        `json:"message"`
	Account       string        `json:"account"`
	SignatureData SignatureData `json:"signature_data"`
}

type SignatureData struct {
	HashedMessage string `json:"hashed_message"`
	R             string `json:"r"`
	S             string `json:"s"`
	V             string `json:"V"`
}

var Queue []VerificationRequest
var Check = make(chan bool)

func main() {

	setupDatabase()
	refreshQueue()
	Check <- true

	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, "https://red-weathered-firefly.avalanche-testnet.quiknode.pro/ext/bc/C/rpc")
	if err != nil {
		log.Fatal(err)
	}
	go handleRequests(client, ctx)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},    // All origins
		AllowedMethods: []string{"POST"}, // Allowing only get, just an example
	})

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/verify", verify).Methods("POST")
	r.Use(mux.CORSMethodMiddleware(r))
	log.Fatal(http.ListenAndServe(":8080", c.Handler(r)))
}

func setupDatabase() {
	if _, err := os.Stat("./sqlite-database.db"); errors.Is(err, os.ErrNotExist) {
		log.Println("Creating sqlite-database.db...")
		file, err := os.Create("sqlite-database.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("sqlite-database.db created")

		db, err := sql.Open("sqlite3", "./sqlite-database.db")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer db.Close()

		log.Println("Create tables...")
		statement, err := db.Prepare(`CREATE TABLE requests (
		"wallet_address" TEXT,		
		"email" TEXT,
		"status" TEXT,
		"message" TEXT,
		"account" TEXT,
		"hashed_message" TEXT,
		"r" TEXT,
		"s" TEXT,
		"v" TEXT
	  );`)
		if err != nil {
			log.Fatal(err.Error())
		}
		_, err = statement.Exec() // Execute SQL Statements
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("Table created")
	}
}

func verify(w http.ResponseWriter, r *http.Request) {
	var requestBody VerificationRequest
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

	err = insertNewRequest(requestBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	Check <- true

	w.Write(res)
}

func handleRequests(client *ethclient.Client, ctx context.Context) {
	for {
		select {
		case <-Check:
			refreshQueue()
			for _, v := range Queue {
				log.Println(v)
				acc := getAccountAuth(client, PrivateKey)
				log.Println(acc)
			}
		}
	}
}

func refreshQueue() {
	var _queue []VerificationRequest

	db, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM requests")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var walletAddress string
		var email string
		var status string
		var message string
		var account string
		var hashed_message string
		var r string
		var s string
		var v string
		row.Scan(&walletAddress, &email, &status, &message, &account, &hashed_message, &r, &s, &v)
		_queue = append(_queue, VerificationRequest{WalletAddress: walletAddress, Email: email, Signature: SignatureStruct{Message: message, Account: account, SignatureData: SignatureData{HashedMessage: hashed_message, R: r, S: s, V: v}}})
	}
	Queue = _queue
}

func insertNewRequest(data VerificationRequest) error {
	db, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	statement, err := db.Prepare(
		`INSERT INTO requests(wallet_address, status, email, hashed_message, r, s, v, message, account) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = statement.Exec(
		data.WalletAddress,
		"pending",
		data.Email,
		data.Signature.SignatureData.HashedMessage,
		data.Signature.SignatureData.R,
		data.Signature.SignatureData.S,
		data.Signature.SignatureData.V,
		data.Signature.Message,
		data.Signature.Account,
	)
	if err != nil {
		return err
	}
	return nil
}

func getAccountAuth(client *ethclient.Client, accountAddress string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	return auth
}