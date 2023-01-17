package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	fmt.Println("Connected to MongoDB Department database")

	collection := client.Database("department_db").Collection("departments")
	mod := mongo.IndexModel{
		Keys: bson.M{
			"did": 1,
		},
		// create UniqueIndex option
		Options: options.Index().SetUnique(true),
	}

	collection.Indexes().CreateOne(ctx, mod)

	mod = mongo.IndexModel{
		Keys: bson.M{
			"eiddpt.eid":   1, // index in ascendingg order
			"projects.pid": 1,
		},
		// create UniqueIndex option
		Options: options.Index().SetUnique(true),
	}

	collection.Indexes().CreateOne(ctx, mod)

	return client
}
