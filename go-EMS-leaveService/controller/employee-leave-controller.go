package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yashjindal28/go-EMS-leaveService/model"
	"github.com/yashjindal28/go-EMS-leaveService/service"
)

func DisplayEmployeeDataEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)
	managerID := params["eid"]

	employeeLeaves, err, err2 := service.DisplayEmployeeLeaves(managerID)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(response).Encode(employeeLeaves)
}

func AddLeaveDataForNewEmployeeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var newEmployeeLeaveData model.EmployeeLeave
	json.NewDecoder(request.Body).Decode(&newEmployeeLeaveData)

	result := service.AddLeaveDataForNewEmployeeByID(newEmployeeLeaveData)

	json.NewEncoder(response).Encode(result)
}

func DeleteLeaveDataForNewEmployeeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)

	childEid := params["childEid"]

	result := service.DeleteEmployeeByID(childEid)

	json.NewEncoder(response).Encode(result)
	// call service to create user in database

}

func ApproveLeaveRequestEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)
	eid := params["childEid"] // childEid is the id for employee under the manager whose leave needs to be aprroved.

	result, err := service.ApproveLeaveRequest(eid)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(response).Encode(result)
}

func RejectLeaveRequestEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)
	eid := params["childEid"]

	result, err := service.RejectLeaveRequest(eid)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(response).Encode(result)
}

func GetLeaveDataForOneEmployeeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)
	eid := params["eid"]

	leaveData := service.GetEmployeeLeaveDataByID(eid)
	json.NewEncoder(response).Encode(leaveData)
}

func RequestLeaveEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var employeeLeaveData model.EmployeeLeave

	json.NewDecoder(request.Body).Decode(&employeeLeaveData)

	result := service.RequestLeaveByID(employeeLeaveData)
	json.NewEncoder(response).Encode(result)
}

func EditEmployeeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)
	eid := params["childEid"]
	var employee model.EmployeeLeave
	json.NewDecoder(request.Body).Decode(&employee)

	result, err := service.UpdateEmployeeByIdUnderManager(eid, employee)

	if err != nil {
		panic(err)
	}
	json.NewEncoder(response).Encode(result)

}
