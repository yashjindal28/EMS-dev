package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/yashjindal28/go-EMS-AuthService/controller"
)

func main() {
	fmt.Println("Employee Management System Auth Service Started ....")

	router := mux.NewRouter()

	router.HandleFunc("/auth/login", controller.LoginEndpoint).Methods("POST")
	router.HandleFunc("/auth/verify", controller.VerifyEndpoint).Methods("GET")
	router.HandleFunc("/auth/create-user/{eid}", controller.CreateUserEndpoint).Methods("POST").Name("CreateUser")
	router.HandleFunc("/auth/{eid}/delete-user/{childEid}", controller.DeleteUserEndpoint).Methods("DELETE").Name("DeleteUser")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
}
