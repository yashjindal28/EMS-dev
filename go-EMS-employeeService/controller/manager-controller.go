// Package classification of Manager API
//
// Documentation for Manager API
//
// Schemes: http
// Basepath: /manager
// Version: 1.0.0
//
// Consumes:
//  - application/json
//
// Produces:
//  - application/json
// swagger:meta

package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/yashjindal28/go-EMS-employeeService/model"
	"github.com/yashjindal28/go-EMS-employeeService/service"
)

func EditEmployeeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)
	eid := params["childEid"]
	var employee model.Employee
	json.NewDecoder(request.Body).Decode(&employee)

	result, err := service.UpdateEmployeeByIdUnderManager(eid, employee)

	if err != nil {
		panic(err)
	}
	json.NewEncoder(response).Encode(result)

}

func DeleteEmployeeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	eid := params["childEid"]

	result := service.DeleteEmployeeByID(eid)
	json.NewEncoder(response).Encode(result)
}

func CreateEmployeeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var employee model.Employee
	json.NewDecoder(request.Body).Decode(&employee)

	result := service.CreateEmployeeUnderManager(employee)

	json.NewEncoder(response).Encode(result)

}

func ListEmployeeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)
	eid := params["eid"]

	employees, err, err2 := service.ListEmployeesUnderManagerById(eid)

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

	json.NewEncoder(response).Encode(employees)
}
