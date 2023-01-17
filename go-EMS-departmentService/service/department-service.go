package service

import (
	"context"
	"fmt"
	"time"

	"github.com/yashjindal28/go-EMS-departmentService/model"
	"github.com/yashjindal28/go-EMS-departmentService/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllDepartmentData() (err error, err2 error, departments []model.Departments) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := repository.GetAllDepartmentData()
	for cursor.Next(ctx) {
		var department model.Departments
		cursor.Decode(&department)
		departments = append(departments, department)
	}
	err2 = cursor.Err()
	if err2 != nil {
		panic(err2)
	}
	return err, err2, departments
}

func AddNewEmployeeToDepartment(employee model.Employee) {

	repository.AddNewEmployeeToDepartment(employee)
}

func DeleteEmployeeByID(childEid string, did string) *mongo.UpdateResult {
	result := repository.DeleteEmployeeByID(childEid, did)

	return result
}

func GetDepartmentByID(did string) (department model.Departments) {

	repository.GetDepartmentByID(did).Decode(&department)
	return department
}

func SaveProjectInDepartment(did string, project model.Projects) {

	repository.SaveProjectInDepartment(did, project)
}

func GetEmployeeDetailsUnderDepartment(did string) []model.Employee {

	var employees []model.Employee

	department, err := repository.GetEmployeeDetailsUnderDepartment(did)
	if err != nil {
		panic(err)
	}
	fmt.Println(department)
	return employees
}

func SaveEmployeeToAProject(customObj model.CustomObj) {

	repository.SaveEmployeeToAProject(customObj)
}

func GetCount() model.Count {
	count := repository.GetCount()

	return count
}

func UpdateEmployeeByIdUnderManager(eid string, employee model.Employee) (result *mongo.UpdateResult, err error) {

	result, err = repository.UpdateEmployeeByIdUnderManager(eid, employee)

	return result, err
}
