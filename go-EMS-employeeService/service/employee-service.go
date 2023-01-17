package service

import (
	"github.com/yashjindal28/go-EMS-employeeService/model"
	"github.com/yashjindal28/go-EMS-employeeService/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPersonalDataOfEmployeeByID(eid string) (personalInfo model.PersonalInfo, err error) {

	personalInfo, err = repository.GetPersonalDataOfEmployeeByID(eid)

	return personalInfo, err

}

func UpdatePersonalInfoByID(eid string, personalInfo model.PersonalInfo) (result *mongo.UpdateResult, err error) {

	address := personalInfo.Address
	contact := personalInfo.Contact

	filter := bson.M{"eid": eid}
	update := bson.M{
		"$set": bson.M{
			"address": address,
			"contact": contact,
		},
	}

	result, err = repository.UpdatePersonalInfoByID(filter, update)

	return result, err
}

func GetOneEmployeeDetailByID(eid string) (employee model.Employee, err error) {

	employee, err = repository.GetOneEmployeeDetailByID(eid)
	return employee, err
}
