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
	"puffinverificationbackend/src/pkg/global"
	"time"
)

func InsertRequest(req global.VerificationRequest, coll string, status string) (primitive.ObjectID, error) {
	req.Status = status
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://verificationBackend:BangFestina@35.84.247.233:27017/PuffinTestnet"))
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
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://verificationBackend:BangFestina@35.84.247.233:27017/PuffinTestnet"))
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