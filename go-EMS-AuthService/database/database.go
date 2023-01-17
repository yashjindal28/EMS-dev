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
	//var data struct {
	//	MONGO_URL string
	//}

	//ctx = context.WithValue(ctx, data.MONGO_URL, os.Getenv("MONGO_URL"))
	//uri := fmt.Sprintf(`mongodb://%s`,
	//	ctx.Value(data.MONGO_URL).(string),
	//)
	//fmt.Println(uri)

	//client, _ = mongo.Connect(ctx, options.Client().ApplyURI(uri))

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	fmt.Println("Connected to MongoDB Auth Database")

	collection1 := client.Database("users").Collection("userAuth")
	collection2 := client.Database("users").Collection("userInfo")
	mod := mongo.IndexModel{
		Keys: bson.M{
			"eid": 1, // index in ascendingg order
		},
		// create UniqueIndex option
		Options: options.Index().SetUnique(true),
	}
	collection1.Indexes().CreateOne(ctx, mod)
	collection2.Indexes().CreateOne(ctx, mod)

	return client
}
