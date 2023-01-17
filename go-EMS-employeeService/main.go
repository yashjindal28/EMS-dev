package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/yashjindal28/go-EMS-employeeService/controller"
)

func main() {
	fmt.Println("Employee Management System")

	//client = ConnectDB()

	router := mux.NewRouter()

	router.HandleFunc("/manager/{eid}/create-employee", controller.CreateEmployeeEndpoint).Methods("POST").Name("CreateEmployee")
	router.HandleFunc("/manager/{eid}/view-employees", controller.ListEmployeeEndpoint).Methods("GET").Name("GetAllEmployeesUnderManager")
	router.HandleFunc("/manager/{eid}/employee-detail/{childEid}", controller.GetOneEmployeeEndpoint).Methods("GET").Name("GetOneEmployeeDetails")
	router.HandleFunc("/manager/{eid}/delete-employee/{childEid}", controller.DeleteEmployeeEndpoint).Methods("DELETE").Name("DeleteEmployee")
	router.HandleFunc("/manager/{eid}/update-employee/{childEid}", controller.EditEmployeeEndpoint).Methods("PUT").Name("UpdateEmployee")

	router.HandleFunc("/employee/{eid}", controller.GetPersonalDataOfEmployeeEndpoint).Methods("GET").Name("GetEmployeePersonalDataByID")
	router.HandleFunc("/employee/{eid}", controller.UpdatePersonalInfoEndpoint).Methods("PUT").Name("UpdatePersonalInfo")

	router.HandleFunc("/allemployees/{eid}", controller.AllEmployeeEndpoint).Methods("GET").Name("AllEmployeesInCompany")
	router.HandleFunc("/allemployees/{eid}/search/{text}", controller.SearchEndpoint).Methods("GET").Name("Search")
	router.HandleFunc("/allemployees/{eid}/filter", controller.FilterEndpoint).Methods("POST").Name("Filter")

	router.Use(controller.AuthorizationHandler)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":84", handlers.CORS(headers, methods, origins)(router)))

}
