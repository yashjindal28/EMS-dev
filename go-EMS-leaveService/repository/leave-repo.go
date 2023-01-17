package repository

import (
	"context"
	"log"
	"time"

	"github.com/yashjindal28/go-EMS-leaveService/model"

	"github.com/yashjindal28/go-EMS-leaveService/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client = database.ConnectDB()

func FindAllEmployeesLeaveData(managerID string) (cursor *mongo.Cursor, err error) {

	collection := client.Database("leave_db").Collection("employeeLeaves")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err = collection.Find(ctx, bson.M{"managerID": managerID})
	//defer cursor.Close(ctx)
	return cursor, err
}

func AddLeaveDataForNewEmployeeByID(newEmployeeLeaveData model.EmployeeLeave) *mongo.InsertOneResult {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("leave_db").Collection("employeeLeaves")
	result, _ := collection.InsertOne(ctx, newEmployeeLeaveData)

	return result
}

func DeleteEmployeeByID(childEid string) *mongo.DeleteResult {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := client.Database("leave_db").Collection("employeeLeaves")
	result, err := collection.DeleteOne(ctx, bson.M{"eid": childEid})
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func RejectLeave(eid string) (result *mongo.UpdateResult, err error) {

	collection := client.Database("leave_db").Collection("employeeLeaves")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"eid": eid}
	//fmt.Println(filter)
	update := bson.M{
		"$set": bson.M{
			"pendingStatus": 0,
			"approved":      0,
		},
	}

	result, err = collection.UpdateMany(ctx, filter, update)
	return result, err
}

func ApproveLeave(eid string, leavesRem int, daysReq int) (result *mongo.UpdateResult, err error) {
	collection := client.Database("leave_db").Collection("employeeLeaves")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"eid": eid}
	//fmt.Println(filter)
	leavesRem = leavesRem - daysReq
	update := bson.M{
		"$set": bson.M{
			"pendingStatus": 0,
			"leavesRem":     leavesRem,
			"approved":      1,
		},
	}

	result, err = collection.UpdateMany(ctx, filter, update)
	return result, err
}

func GetLeaveDataByID(eid string) (leaveData *mongo.SingleResult) {
	collection := client.Database("leave_db").Collection("employeeLeaves")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	leaveData = collection.FindOne(ctx, bson.M{"eid": eid})
	return leaveData
}

func RequestLeaveByID(employeeLeaveData model.EmployeeLeave) *mongo.UpdateResult {

	collection := client.Database("leave_db").Collection("employeeLeaves")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{"eid", employeeLeaveData.EmployeeID}}
	replacement := employeeLeaveData
	result, err := collection.ReplaceOne(ctx, filter, replacement)
	if err != nil {
		panic(err)
	}
	return result
}

func UpdateEmployeeByIdUnderManager(eid string, employee model.EmployeeLeave) (result *mongo.UpdateResult, err error) {

	collection := client.Database("leave_db").Collection("employeeLeaves")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.D{{"eid", eid}}
	update := bson.M{
		"$set": bson.M{
			"firstname": employee.Firstname,
			"lastname":  employee.Lastname,
			"manager":   employee.Manager,
			"managerID": employee.ManagerID,
		},
	}
	_, err = collection.UpdateOne(ctx, filter, update) // you can simply replace using replace one command and decoded personalInfo obejct
	if err != nil {
		panic(err)
	}

	filter = bson.D{{"managerID", eid}}
	update = bson.M{
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
