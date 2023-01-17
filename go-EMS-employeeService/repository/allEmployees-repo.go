package repository

import (
	"context"
	"time"

	"github.com/yashjindal28/go-EMS-employeeService/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client = database.ConnectDB()

func FindAllEmployees() (cursor *mongo.Cursor, err error) {

	collection := client.Database("employees_db").Collection("employees")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err = collection.Find(ctx, bson.M{})

	return cursor, err

}

func Search(text string) (cursor *mongo.Cursor, err error) {
	collection := client.Database("employees_db").Collection("employees")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.D{{"$text", bson.D{{"$search", text}}}}
	cursor, err = collection.Find(ctx, filter)

	return cursor, err
}

func Filter(filter bson.M) (cursor *mongo.Cursor, err error) {

	collection := client.Database("employees_db").Collection("employees")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err = collection.Find(ctx, filter)

	return cursor, err
}
