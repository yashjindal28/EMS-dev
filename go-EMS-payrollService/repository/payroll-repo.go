package repository

import (
	"context"
	"log"
	"time"

	"github.com/yashjindal28/go-EMS-payrollService/database"
	"github.com/yashjindal28/go-EMS-payrollService/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client = database.ConnectDB()

func FindAllPayrollData() (cursor *mongo.Cursor, err error) {

	collection := client.Database("payroll_db").Collection("payrollInfo")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err = collection.Find(ctx, bson.M{})

	return cursor, err
}

func GetPayrollInfoOfEmployeeByID(eid string) (result *mongo.SingleResult) {

	collection := client.Database("payroll_db").Collection("payrollInfo")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result = collection.FindOne(ctx, bson.M{"eid": eid})
	return result
}

func UpdatePayrollInfo(eid string) {

	collection := client.Database("payroll_db").Collection("payrollInfo")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"eid": eid}
	update := bson.M{
		"$set": bson.M{
			"isCheckIssued": 1,
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update) // you can simply replace using replace one command and decoded personalInfo obejct
	if err != nil {
		panic(err)
	}
}

func AddPayrollInfo(payInfo model.PayrollData) *mongo.InsertOneResult {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	payInfo.IsCheckIssued = 0
	collection5 := client.Database("payroll_db").Collection("payrollInfo")
	result, _ := collection5.InsertOne(ctx, payInfo)

	return result
}

func DeleteEmployeeByID(childEid string) *mongo.DeleteResult {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection5 := client.Database("payroll_db").Collection("payrollInfo")
	result, err := collection5.DeleteOne(ctx, bson.M{"eid": childEid})
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func UpdateEmployeeByIdUnderManager(eid string, employee model.PayrollData) (result *mongo.UpdateResult, err error) {

	collection := client.Database("payroll_db").Collection("payrollInfo")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.D{{"eid", eid}}
	update := bson.M{
		"$set": bson.M{
			"firstname": employee.Firstname,
			"lastname":  employee.Lastname,
			"desg":      employee.Designation,
			"did":       employee.DepartmentID,
			"dpt":       employee.Department,
			"manager":   employee.Manager,
			"managerID": employee.ManagerID,
			"salary":    employee.Salary,
			"email":     employee.Email,
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
