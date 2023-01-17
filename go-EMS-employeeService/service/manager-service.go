package service

import (
	"context"
	"time"

	"github.com/yashjindal28/go-EMS-employeeService/model"
	"github.com/yashjindal28/go-EMS-employeeService/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func ListEmployeesUnderManagerById(eid string) (employees []model.Employee, err error, err2 error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := repository.ListEmployeesUnderManagerById(eid)

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var employee model.Employee
		cursor.Decode(&employee)
		employees = append(employees, employee)
	}
	err2 = cursor.Err()

	return employees, err, err2
}

func UpdateEmployeeByIdUnderManager(eid string, employee model.Employee) (result *mongo.UpdateResult, err error) {

	result, err = repository.UpdateEmployeeByIdUnderManager(eid, employee)

	return result, err
}

func DeleteEmployeeByID(eid string) *mongo.DeleteResult {
	result := repository.DeleteEmployeeByID(eid)

	return result
}

func CreateEmployeeUnderManager(employee model.Employee) *mongo.InsertOneResult {
	result := repository.CreateEmployeeUnderManager(employee)

	return result
}
