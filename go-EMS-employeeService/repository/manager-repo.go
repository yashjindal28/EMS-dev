package repository

import (
	"context"
	"log"
	"time"

	"github.com/yashjindal28/go-EMS-employeeService/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ListEmployeesUnderManagerById(eid string) (cursor *mongo.Cursor, err error) {

	collection := client.Database("employees_db").Collection("employees")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err = collection.Find(ctx, bson.M{"managerID": eid})

	return cursor, err
}

func UpdateEmployeeByIdUnderManager(eid string, employee model.Employee) (result *mongo.UpdateResult, err error) {

	collection := client.Database("employees_db").Collection("employees")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{"eid", eid}}
	replacement := employee
	result, err = collection.ReplaceOne(ctx, filter, replacement)

	filter = bson.D{{"managerID", eid}}
	update := bson.M{
		"$set": bson.M{
			"manager": employee.Firstname + " " + employee.Lastname,
		},
	}
	_, err = collection.UpdateMany(ctx, filter, update) // you can simply replace using replace one command and decoded personalInfo obejct
	if err != nil {
		panic(err)
	}

	return result, err
}

func DeleteEmployeeByID(eid string) *mongo.DeleteResult {

	collection := client.Database("employees_db").Collection("employees") // have to delete personal , leave info too
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.DeleteOne(ctx, bson.M{"eid": eid})
	if err != nil {
		log.Fatal(err)
	}
	collection2 := client.Database("employees_db").Collection("personalInfo")
	result, err = collection2.DeleteOne(ctx, bson.M{"eid": eid})
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func CreateEmployeeUnderManager(employee model.Employee) *mongo.InsertOneResult {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := client.Database("employees_db").Collection("employees")
	result, _ := collection.InsertOne(ctx, employee)

	var personalInfo = model.PersonalInfo{EmployeeID: employee.EmployeeID, Address: "", Contact: ""}
	collection2 := client.Database("employees_db").Collection("personalInfo")
	collection2.InsertOne(ctx, personalInfo)

	return result
}
