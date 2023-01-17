package controller

import (
	"encoding/json"
	"net/http"

	"github.com/yashjindal28/go-EMS-payrollService/model"

	"github.com/gorilla/mux"
	"github.com/yashjindal28/go-EMS-payrollService/service"
)

func AllEmployeeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	err, err2, payrollData := service.FindAll()

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

	json.NewEncoder(response).Encode(payrollData)

}
func IssuePaycheckEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)
	eid := params["childEid"]
	//fmt.Println(eid)
	service.CalSalaryAndMail(eid)

	//json.NewEncoder(response).Encode(employees)
}

func AddPayrollInfoEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var payInfo model.PayrollData
	json.NewDecoder(request.Body).Decode(&payInfo)
	result := service.AddPayrollInfo(payInfo)

	json.NewEncoder(response).Encode(result)

}

func DeletePayrollInfoForEmployeeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)

	childEid := params["childEid"]

	result := service.DeleteEmployeeByID(childEid)

	json.NewEncoder(response).Encode(result)
	// call service to create user in database

}

func EditEmployeeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)
	eid := params["childEid"]
	var employee model.PayrollData
	json.NewDecoder(request.Body).Decode(&employee)

	result, err := service.UpdateEmployeeByIdUnderManager(eid, employee)

	if err != nil {
		panic(err)
	}
	json.NewEncoder(response).Encode(result)

}
