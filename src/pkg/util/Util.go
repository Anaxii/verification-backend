package util

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func GetOID(id string) (primitive.ObjectID, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
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
