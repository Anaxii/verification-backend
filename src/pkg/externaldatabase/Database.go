package externaldatabase

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"puffinverificationbackend/src/pkg/config"
	"puffinverificationbackend/src/pkg/global"
	"time"
)

func InsertRequest(req global.VerificationRequest, coll string, status string) (primitive.ObjectID, error) {
	req.Status = status
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoDBURI))
	if err != nil {
		log.Println(err)
		return primitive.ObjectID{}, err
	}
	defer client.Disconnect(ctx)

	requestsCollection := client.Database("PuffinTestnet").Collection(coll)

	insertResult, err := requestsCollection.InsertOne(ctx, req)
	if err != nil {
		log.Println(err)
		return primitive.ObjectID{}, err
	}
	return insertResult.InsertedID.(primitive.ObjectID), nil
}

func DenyRequest(req global.VerificationRequest, reason string) error {
	return updateRequest(req, "denied", reason)
}

func ApproveRequest(req global.VerificationRequest) error {
	return updateRequest(req, "approved", "approved")
}

func updateRequest(req global.VerificationRequest, coll string, status string) error {
	oid, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoDBURI))
	if err != nil {
		log.Println(err)
		return err
	}
	defer client.Disconnect(ctx)

	requestsCollection := client.Database("PuffinTestnet").Collection("requests")
	request := requestsCollection.FindOne(context.TODO(), bson.D{{"_id", oid}})
	var result global.VerificationRequest
	err = request.Decode(&result)
	if err != nil {
		return err
	}

	_, err = InsertRequest(result, coll, status)
	if err == nil {
		requestsCollection.DeleteOne(context.TODO(), bson.D{{"_id", oid}})
		return nil
	}

	return errors.New("failed to remove request")
}

func CheckIfExists(walletAddress string, table string) bool {
	log.Println(config.MongoDBURI)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoDBURI))
	if err != nil {
		log.Println("checkifexist", err)
		return false
	}
	defer client.Disconnect(ctx)

	requestsCollection := client.Database("PuffinTestnet").Collection(table)
	request := requestsCollection.FindOne(context.TODO(), bson.D{{"wallet_address", walletAddress}})
	var result global.VerificationRequest
	err = request.Decode(&result)
	if err != nil {
		log.Println("checkifexist", err)
		return false
	}
	return true
}
