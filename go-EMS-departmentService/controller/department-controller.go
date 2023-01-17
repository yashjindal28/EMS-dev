package controller

import (
	"encoding/json"
	"net/http"

	"github.com/yashjindal28/go-EMS-departmentService/model"

	"github.com/gorilla/mux"
	"github.com/yashjindal28/go-EMS-departmentService/service"
)

func GetDepartmentDataEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	err, err2, departments := service.GetAllDepartmentData()

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

	json.NewEncoder(response).Encode(departments)
}

func AddNewEmployeeToDepartmentEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var employee model.Employee

	json.NewDecoder(request.Body).Decode(&employee)
	service.AddNewEmployeeToDepartment(employee)

	//json.NewEncoder(response).Encode(department)
}

func DeleteEmployeeFromDepartmentEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)

	childEid := params["childEid"]
	did := params["did"]

	result := service.DeleteEmployeeByID(childEid, did)

	json.NewEncoder(response).Encode(result)
	// call service to create user in database

}

func GetDepartmentByIDEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)
	did, _ := params["did"]
	department := service.GetDepartmentByID(did)

	json.NewEncoder(response).Encode(department)
}

func SaveProjectEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var project model.Projects
	params := mux.Vars(request)
	did, _ := params["did"]
	json.NewDecoder(request.Body).Decode(&project)
	service.SaveProjectInDepartment(did, project)

	//json.NewEncoder(response).Encode(department)
}

func GetEmployeeDetailsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)
	did, _ := params["did"]
	employees := service.GetEmployeeDetailsUnderDepartment(did)

	json.NewEncoder(response).Encode(employees)
}

func SaveEmployeeToAProjectEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var customObj model.CustomObj

	json.NewDecoder(request.Body).Decode(&customObj)
	service.SaveEmployeeToAProject(customObj)

	//json.NewEncoder(response).Encode(employees)
}

func GetCountEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	count := service.GetCount()

	json.NewEncoder(response).Encode(count)
}

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
