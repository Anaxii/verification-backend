package util

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func GetOID(id string) (primitive.ObjectID, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return primitive.ObjectID{}, err
	}
	return oid, err
}
