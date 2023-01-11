package externaldatabase

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"puffinverificationbackend/internal/config"
	"puffinverificationbackend/internal/global"
	"puffinverificationbackend/pkg/util"
	"time"
)

func InsertRequest(req interface{}, coll string) (primitive.ObjectID, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoDBURI))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Database:InsertRequest"}).Error("Failed to connect to mongodb client")
		return primitive.ObjectID{}, err
	}
	defer client.Disconnect(ctx)

	requestsCollection := client.Database("puffin").Collection(coll)

	insertResult, err := requestsCollection.InsertOne(ctx, req)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Database:InsertRequest", "collection": coll, "request_data": req}).Error("Failed to insert into collection")
		return primitive.ObjectID{}, err
	}
	return insertResult.InsertedID.(primitive.ObjectID), nil
}

func DenyRequest(req global.AccountRequest, reason string, coll string) error {
	if oid, err := util.GetOID(req.ID); err == nil {
		return updateRequest(coll, coll, oid, "denied")
	}
	return errors.New("could not get oid")
}

func ApproveRequest(req global.AccountRequest, coll string) error {
	if oid, err := util.GetOID(req.ID); err == nil {
		err = updateRequest(coll, coll, oid, "approved")
		if err != nil {
			return err
		}
		req.Status = "approved"
		_, err = InsertRequest(req, "accounts")
		return err
	}
	return errors.New("could not get oid")
}

func DenySubRequest(req global.SubAccountRequest, reason string, coll string) error {
	if oid, err := util.GetOID(req.ID); err == nil {
		return updateRequest(coll, coll, oid, "denied")
	}
	return errors.New("could not get oid")
}

func ApproveSubRequest(req global.SubAccountRequest, coll string) error {
	if oid, err := util.GetOID(req.ID); err == nil {
		err = updateRequest(coll, coll, oid, "approved")
		if err != nil {
			return err
		}
		req.Status = "approved"
		_, err = InsertRequest(req, "subaccounts")
		return err
	}
	return errors.New("could not get oid")
}

func updateRequest(collection string, coll string, oid  primitive.ObjectID, status string) error {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoDBURI))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Database:updateRequest"}).Error("Failed to connect to mongodb client")
		return err
	}
	defer client.Disconnect(ctx)

	requestsCollection := client.Database("PuffinTestnet").Collection(collection)
	filter := bson.D{{"_id", oid}}
	update := bson.D{{"$set", bson.D{{"status", status}}}}
	_, err = requestsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func CheckIfExists(walletAddress string, table string, key string) (bool, global.AccountRequest) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoDBURI))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Database:CheckIfExists"}).Error("Failed to connect to mongodb client")
		return false, global.AccountRequest{}
	}
	defer client.Disconnect(ctx)

	requestsCollection := client.Database("puffin").Collection(table)
	request := requestsCollection.FindOne(context.TODO(), bson.D{{key, walletAddress}})
	var result global.AccountRequest
	err = request.Decode(&result)
	if err != nil {
		//log.WithFields(log.Fields{"error": err.Error(), "file": "Database:CheckIfExists", "table": table, "key": key, "value": walletAddress}).Warn("Failed to decode verification results")
		return false, global.AccountRequest{}
	}
	return true, result
}

func GetCountries() global.Countries {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoDBURI))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Database:CheckIfExists"}).Error("Failed to connect to mongodb client")
		return global.Countries{}
	}
	defer client.Disconnect(ctx)

	requestsCollection := client.Database("PuffinTestnet").Collection("countries")
	request := requestsCollection.FindOne(context.TODO(), bson.D{})
	var result global.Countries
	err = request.Decode(&result)
	if err != nil {
		return global.Countries{}
	}
	return result
}

func GetAllUsers() ([]global.AccountRequest, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoDBURI))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Database:CheckIfExists"}).Error("Failed to connect to mongodb client")
		return []global.AccountRequest{}, err
	}
	defer client.Disconnect(ctx)

	requestsCollection := client.Database("PuffinTestnet").Collection("accounts")
	cur, err := requestsCollection.Find(context.TODO(), bson.D{})
	if err !=nil {
		log.Error(err)
	}

	var results []global.AccountRequest
	for cur.Next(context.TODO()) {
		var result global.AccountRequest
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, result)

	}

	return results, nil
}
