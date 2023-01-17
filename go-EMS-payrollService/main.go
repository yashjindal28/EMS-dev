package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/yashjindal28/go-EMS-payrollService/controller"
)

func main() {
	fmt.Println("Employee Management System Payroll Service")

	router := mux.NewRouter()

	router.HandleFunc("/payroll/{eid}", controller.AllEmployeeEndpoint).Methods("GET").Name("GetPayrollDataForAllEmployess")
	router.HandleFunc("/payroll/{eid}/add-info", controller.AddPayrollInfoEndpoint).Methods("POST").Name("AddPayrollInfo")
	router.HandleFunc("/payroll/{eid}/delete-info/{childEid}", controller.DeletePayrollInfoForEmployeeEndpoint).Methods("DELETE").Name("DeleteEmployee")
	router.HandleFunc("/payroll/{eid}/update-employee/{childEid}", controller.EditEmployeeEndpoint).Methods("PUT").Name("UpdateEmployee")
	router.HandleFunc("/payroll/{eid}/issue/{childEid}", controller.IssuePaycheckEndpoint).Methods("GET").Name("IssuePayCheckByEmployeeID")
	router.Use(controller.AuthorizationHandler)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":81", handlers.CORS(headers, methods, origins)(router)))
}
