package service

import (
	"context"
	"time"

	"github.com/yashjindal28/go-EMS-leaveService/model"
	"github.com/yashjindal28/go-EMS-leaveService/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func DisplayEmployeeLeaves(managerID string) (EmployeeLeaves []model.EmployeeLeave, err error, err2 error) {

	//var EmployeeLeaves []EmployeeLeave
	cursor, err := repository.FindAllEmployeesLeaveData(managerID)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	for cursor.Next(ctx) {
		var employeeLeave model.EmployeeLeave
		cursor.Decode(&employeeLeave)
		EmployeeLeaves = append(EmployeeLeaves, employeeLeave)
	}
	err2 = cursor.Err()

	return EmployeeLeaves, err, err2

}

func AddLeaveDataForNewEmployeeByID(newEmployeeLeaveData model.EmployeeLeave) *mongo.InsertOneResult {

	newEmployeeLeaveData.LeavesRemaining = 20
	newEmployeeLeaveData.DaysRequested = 0
	newEmployeeLeaveData.PendingStatus = 0
	newEmployeeLeaveData.Approved = 0

	result := repository.AddLeaveDataForNewEmployeeByID(newEmployeeLeaveData)

	return result
}

func DeleteEmployeeByID(childEid string) *mongo.DeleteResult {
	result := repository.DeleteEmployeeByID(childEid)

	return result
}

func RejectLeaveRequest(eid string) (*mongo.UpdateResult, error) {

	result, err := repository.RejectLeave(eid)
	return result, err
}

func ApproveLeaveRequest(eid string) (*mongo.UpdateResult, error) {

	var leaveData model.EmployeeLeave
	repository.GetLeaveDataByID(eid).Decode(&leaveData)

	result, err := repository.ApproveLeave(eid, leaveData.LeavesRemaining, leaveData.DaysRequested)
	return result, err
}

func GetEmployeeLeaveDataByID(eid string) (leaveData model.EmployeeLeave) {

	repository.GetLeaveDataByID(eid).Decode(&leaveData)
	return leaveData
}

func RequestLeaveByID(employeeLeaveData model.EmployeeLeave) *mongo.UpdateResult {

	employeeLeaveData.PendingStatus = 1
	employeeLeaveData.Approved = 0
	result := repository.RequestLeaveByID(employeeLeaveData)
	return result
}

func UpdateEmployeeByIdUnderManager(eid string, employee model.EmployeeLeave) (result *mongo.UpdateResult, err error) {

	result, err = repository.UpdateEmployeeByIdUnderManager(eid, employee)

	return result, err
}
