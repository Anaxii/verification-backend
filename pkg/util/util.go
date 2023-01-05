package util

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"net/http"
)

func GetOID(id string) (primitive.ObjectID, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Util:GetOID"}).Error("Failed to convert id string to primitive.ObjectID")
		return primitive.ObjectID{}, err
	}
	return oid, err
}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}


func RandStringRunes(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
