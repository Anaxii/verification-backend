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
	"puffinverificationbackend/src/pkg/util"
	"time"
)

func InsertRequest(req interface{}, coll string) (primitive.ObjectID, error) {
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

func DenyRequest(req global.VerificationRequest, reason string, coll string) error {
	if oid, err := util.GetOID(req.ID); err == nil {
		return updateRequest(coll, "denied", oid)
	}
	return errors.New("could not get oid")
}

func ApproveRequest(req global.VerificationRequest, coll string) error {
	if oid, err := util.GetOID(req.ID); err == nil {
		return updateRequest(coll, "approved", oid)
	}
	return errors.New("could not get oid")
}

func DenySubRequest(req global.SubAccountRequest, reason string, coll string) error {
	if oid, err := util.GetOID(req.ID); err == nil {
		return updateRequest(coll, "denied_subaccounts", oid)
	}
	return errors.New("could not get oid")
}

func ApproveSubRequest(req global.SubAccountRequest, coll string) error {
	if oid, err := util.GetOID(req.ID); err == nil {
		return updateRequest(coll, "subaccounts", oid)
	}
	return errors.New("could not get oid")
}

func updateRequest(collection string, coll string, oid  primitive.ObjectID) error {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoDBURI))
	if err != nil {
		log.Println(err)
		return err
	}
	defer client.Disconnect(ctx)

	requestsCollection := client.Database("PuffinTestnet").Collection(collection)
	request := requestsCollection.FindOne(context.TODO(), bson.D{{"_id", oid}})
	var result interface{}
	err = request.Decode(&result)
	if err != nil {
		return err
	}

	_, err = InsertRequest(result, coll)
	if err == nil {
		requestsCollection.DeleteOne(context.TODO(), bson.D{{"_id", oid}})
		return nil
	}

	return errors.New("failed to remove request")
}

func CheckIfExists(walletAddress string, table string, key string) (bool, global.VerificationRequest) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoDBURI))
	if err != nil {
		log.Println("checkifexist", err)
		return false, global.VerificationRequest{}
	}
	defer client.Disconnect(ctx)

	requestsCollection := client.Database("PuffinTestnet").Collection(table)
	request := requestsCollection.FindOne(context.TODO(), bson.D{{key, walletAddress}})
	var result global.VerificationRequest
	err = request.Decode(&result)
	if err != nil {
		return false, global.VerificationRequest{}
	}
	return true, result
}
