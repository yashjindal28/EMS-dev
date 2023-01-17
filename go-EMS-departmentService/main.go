package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/yashjindal28/go-EMS-departmentService/controller"
)

func main() {
	fmt.Println("Employee Management System Department Service")

	router := mux.NewRouter()

	router.HandleFunc("/department/{eid}", controller.GetDepartmentDataEndpoint).Methods("GET").Name("GetAllDepartmentsData")
	router.HandleFunc("/department/{eid}/add-employee", controller.AddNewEmployeeToDepartmentEndpoint).Methods("POST").Name("AddNewEmployeeToDepartment")
	router.HandleFunc("/department/{eid}/delete-employee/{childEid}/{did}", controller.DeleteEmployeeFromDepartmentEndpoint).Methods("DELETE").Name("DeleteEmployee")
	router.HandleFunc("/department/{eid}/one/{did}", controller.GetDepartmentByIDEndpoint).Methods("GET").Name("GetADepartmentByID")
	router.HandleFunc("/department/{eid}/details/{did}", controller.GetEmployeeDetailsEndpoint).Methods("GET").Name("GetEmployeeDetailsOfADepartment")
	router.HandleFunc("/department/{eid}/save-project/{did}", controller.SaveProjectEndpoint).Methods("PUT").Name("SaveProjectInADepartment")
	router.HandleFunc("/department/{eid}/save-employee", controller.SaveEmployeeToAProjectEndpoint).Methods("POST").Name("SaveEmployeeToAProject")
	router.HandleFunc("/department/{eid}/get-count", controller.GetCountEndpoint).Methods("GET").Name("GetCount")
	router.HandleFunc("/department/{eid}/update-employee/{childEid}", controller.EditEmployeeEndpoint).Methods("PUT").Name("UpdateEmployee")

	router.Use(controller.AuthorizationHandler)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":83", handlers.CORS(headers, methods, origins)(router)))
}
