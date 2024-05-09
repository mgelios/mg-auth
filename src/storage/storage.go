package storage

import (
	"context"
	"fmt"
	"log"
	"mg-auth/src/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongoClient() *mongo.Client {
	clientOption := options.Client().ApplyURI("mongodb://localhost:19000")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal(err)
	}
	client.Database("auth").CreateCollection(ctx, "user")
	if err != nil {
		fmt.Println("Error while creating user collection")
	}
	return client
}

// func disconnectMongoClient() {
// 	if err = mongo_client.Disconnect(ctx); err != nil {
// 		panic(err)
// 	}
// }

var mongo_client *mongo.Client = initMongoClient()

func UpdateUser(id string, record model.User) {
	collection := mongo_client.Database("auth").Collection("user")
	_, err := collection.UpdateByID(context.Background(), id, record)
	if err != nil {
		log.Fatal(err)
	}
}

func PutUser(record model.User) {
	collection := mongo_client.Database("auth").Collection("user")
	_, err := collection.InsertOne(context.Background(), record)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserById(id string) (model.User, error) {
	collection := mongo_client.Database("auth").Collection("user")
	result := model.User{}
	filter := bson.D{{"id", id}}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	return result, err
}
