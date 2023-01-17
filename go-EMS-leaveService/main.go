package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/yashjindal28/go-EMS-leaveService/controller"
)

func main() {
	fmt.Println("Employee Management System Leave Service")

	router := mux.NewRouter()

	router.HandleFunc("/leave/employee/{eid}/add-info", controller.AddLeaveDataForNewEmployeeEndpoint).Methods("POST").Name("AddLeaveDataForNewEmployee")
	router.HandleFunc("/leave/employee/{eid}/delete-info/{childEid}", controller.DeleteLeaveDataForNewEmployeeEndpoint).Methods("DELETE").Name("DeleteEmployee")
	router.HandleFunc("/leave/employee/{eid}", controller.GetLeaveDataForOneEmployeeEndpoint).Methods("GET").Name("EmployeeLeaveData")
	router.HandleFunc("/leave/employee/request-leave/{eid}", controller.RequestLeaveEndpoint).Methods("PUT").Name("EmployeeRequestLeave")
	router.HandleFunc("/leave/manager/{eid}", controller.DisplayEmployeeDataEndpoint).Methods("GET").Name("AllEmployeesUnderManagerLeaveData")
	router.HandleFunc("/leave/manager/{eid}/approve/{childEid}", controller.ApproveLeaveRequestEndpoint).Methods("GET").Name("AprroveLeave")
	router.HandleFunc("/leave/manager/{eid}/reject/{childEid}", controller.RejectLeaveRequestEndpoint).Methods("GET").Name("RejectLeave")
	router.HandleFunc("/leave/{eid}/update-employee/{childEid}", controller.EditEmployeeEndpoint).Methods("PUT").Name("UpdateEmployee")

	router.Use(controller.AuthorizationHandler)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":82", handlers.CORS(headers, methods, origins)(router)))
}
