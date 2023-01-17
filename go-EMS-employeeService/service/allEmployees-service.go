package service

import (
	"context"
	"time"

	"github.com/yashjindal28/go-EMS-employeeService/model"
	"github.com/yashjindal28/go-EMS-employeeService/repository"
	"go.mongodb.org/mongo-driver/bson"
)

func FindAllEmployees() (employees []model.Employee, err error, err2 error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := repository.FindAllEmployees()

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var employee model.Employee
		cursor.Decode(&employee)
		employees = append(employees, employee)
	}
	err2 = cursor.Err()

	return employees, err, err2
}

func Search(text string) (employees []model.Employee) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := repository.Search(text)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var employee model.Employee
		cursor.Decode(&employee)
		employees = append(employees, employee)
	}
	err = cursor.Err()
	if err != nil {
		panic(err)
	}

	return employees
}

func Filter(criteria model.Employee) (employees []model.Employee) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	desg := criteria.Designation
	dpt := criteria.Department
	location := criteria.Location

	filter := bson.M{}
	if desg != "" {
		//filter = bson.M{"title": params.Title}
		filter["desg"] = desg
	}
	if dpt != "" {
		//filter = bson.M{"title": params.Title}
		filter["dpt"] = dpt
	}
	if location != "" {
		//filter = bson.M{"title": params.Title}
		filter["location"] = location
	}

	cursor, err := repository.Filter(filter)

	if err != nil {
		panic(err)
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var employee model.Employee
		cursor.Decode(&employee)
		employees = append(employees, employee)
	}
	err = cursor.Err()
	if err != nil {
		panic(err)
	}

	return employees
}
