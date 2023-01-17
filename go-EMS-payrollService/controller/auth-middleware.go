package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yashjindal28/go-EMS-payrollService/service"
)

func AuthorizationHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		currentRoute := mux.CurrentRoute(request)         // Getting the current route
		currentRouteVars := mux.Vars(request)             // Getting the vars of the current route
		authHeader := request.Header.Get("Authorization") // Getting header of the route
		//fmt.Println(authHeader)
		if authHeader != "" {
			token := service.GetTokenFromHeader(authHeader) // Extracting the token from the header
			//fmt.Println(token)
			isAuthorized := service.IsAuthorized(token, currentRoute.GetName(), currentRouteVars) // checking whter we can access the data from the current route or not. like only payroll manager can access the payroll route
			// if authroized then next api or route can be called. else not

			if isAuthorized {
				next.ServeHTTP(response, request) // calling payroll route to siplay all employees
			} else {
				//  sending response to frontend that you can't visit this route
				response.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(response, "Unauthorized")
			}

		} else {
			response.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(response, "missing token")
		}

	})
}
