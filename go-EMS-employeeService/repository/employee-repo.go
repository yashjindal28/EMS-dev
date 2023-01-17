package repository

import (
	"context"
	"time"

	"github.com/yashjindal28/go-EMS-employeeService/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPersonalDataOfEmployeeByID(eid string) (personalInfo model.PersonalInfo, err error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection2 := client.Database("employees_db").Collection("personalInfo")
	err = collection2.FindOne(ctx, model.PersonalInfo{EmployeeID: eid}).Decode(&personalInfo)

	return personalInfo, err

}

func UpdatePersonalInfoByID(filter bson.M, update bson.M) (result *mongo.UpdateResult, err error) {

	collection2 := client.Database("employees_db").Collection("personalInfo")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err = collection2.UpdateMany(ctx, filter, update) // you can simply replace using replace one command and decoded personalInfo obejct
	return result, err
}

func GetOneEmployeeDetailByID(eid string) (employee model.Employee, err error) {

	collection := client.Database("employees_db").Collection("employees")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = collection.FindOne(ctx, model.Employee{EmployeeID: eid}).Decode(&employee)

	return employee, err
}
