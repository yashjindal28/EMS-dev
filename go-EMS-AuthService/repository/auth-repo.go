package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/yashjindal28/go-EMS-AuthService/database"
	"github.com/yashjindal28/go-EMS-AuthService/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client = database.ConnectDB()

func FindBy(loginRequest model.LoginRequest) (*model.Login, error) {
	var login model.Login
	collection := client.Database("users").Collection("userAuth")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	pipeline := []bson.M{
		bson.M{"$match": bson.M{"eid": loginRequest.EmployeeID, "password": loginRequest.Password}},
		bson.M{
			"$lookup": bson.M{
				"from":         "userInfo",
				"localField":   "eid",
				"foreignField": "eid",
				"as":           "info",
			},
		},
		bson.M{"$unwind": "$info"},
		bson.M{"$project": bson.M{
			"password": 0,
			"_id":      0,
		},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		panic(err)
	}
	if cursor.RemainingBatchLength() == 0 { // checking if no result is returned then returing nil value for login and unauthorized err.
		return nil, errors.New("invalid credentials")
	}
	for cursor.Next(ctx) {
		cursor.Decode(&login)
	}
	if err := cursor.Close(ctx); err != nil {
		panic(err)
	}

	return &login, err
}

func CreateUserByID(userInfo model.UserInfo) *mongo.InsertOneResult {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection1 := client.Database("users").Collection("userAuth")
	result, _ := collection1.InsertOne(ctx, bson.M{"eid": userInfo.EmployeeID, "password": userInfo.EmployeeID})

	collection2 := client.Database("users").Collection("userInfo")
	result, _ = collection2.InsertOne(ctx, bson.M{"eid": userInfo.EmployeeID, "desg": userInfo.Designation})

	return result
}

func DeleteUserByID(childEid string) *mongo.DeleteResult {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection1 := client.Database("users").Collection("userAuth")
	result, err := collection1.DeleteOne(ctx, bson.M{"eid": childEid})
	if err != nil {
		log.Fatal(err)
	}

	collection2 := client.Database("users").Collection("userInfo")
	result, err = collection2.DeleteOne(ctx, bson.M{"eid": childEid})
	if err != nil {
		log.Fatal(err)
	}

	return result
}
