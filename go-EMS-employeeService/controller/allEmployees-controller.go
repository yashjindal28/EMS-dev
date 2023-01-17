package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/yashjindal28/go-EMS-employeeService/model"
	"github.com/yashjindal28/go-EMS-employeeService/service"

	"github.com/gorilla/mux"
)

func AllEmployeeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	employees, err, err2 := service.FindAllEmployees()

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

func SearchEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)
	var text string = params["text"]

	employees := service.Search(text)

	json.NewEncoder(response).Encode(employees)
}

func FilterEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var criteria model.Employee
	json.NewDecoder(request.Body).Decode(&criteria)
	employees := service.Filter(criteria)

	json.NewEncoder(response).Encode(employees)

}
