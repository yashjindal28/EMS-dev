package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yashjindal28/go-EMS-AuthService/model"
	"github.com/yashjindal28/go-EMS-AuthService/service"
)

func LoginEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var loginRequest model.LoginRequest
	err := json.NewDecoder(request.Body).Decode(&loginRequest)
	if err != nil {
		fmt.Println("Error while decoding login request: " + err.Error())
		response.WriteHeader(http.StatusBadRequest)
	} else {
		token, desg, err := service.Login(loginRequest)
		if err != nil {
			writeResponse(response, http.StatusUnauthorized, notAuthorizedResponse(" Invalid Credentials "))
		} else {
			writeResponse(response, http.StatusOK, map[string]string{"jwtToken": *token, "desg": desg})
		}
	}

	//json.NewEncoder(response).Encode(result)

}

/*
  Sample URL string for calling verify api by other services like payroll , etc
 http://localhost:8080/auth/verify?token=somevalidtokenstring&routeName=GetCustomer&eid=2000&desg=95470
*/

func VerifyEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	urlParams := make(map[string]string)

	// converting Query in request to map type to extract all the query parameters
	for k := range request.URL.Query() {
		urlParams[k] = request.URL.Query().Get(k)
	}

	if urlParams["token"] != "" {
		// if there is token present then verify token and send response
		appErr := service.Verify(urlParams)
		if appErr != nil {
			writeResponse(response, http.StatusUnauthorized, notAuthorizedResponse(" Token Verification Failed "))
			panic(appErr)

		} else {

			writeResponse(response, http.StatusOK, authorizedResponse())
		}
	} else {
		// if there is NO token present then verify token and send response missing token
		//fmt.Println("Missing token")
		writeResponse(response, http.StatusForbidden, notAuthorizedResponse("missing token"))
	}
	//json.NewEncoder(response).Encode(result)

}

func CreateUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)
	eid := params["eid"] // this is the manager id who wants to create a employee under him which we need to verify with the token
	var userInfo model.UserInfo
	json.NewDecoder(request.Body).Decode(&userInfo)

	authHeader := request.Header.Get("Authorization")
	currentRoute := mux.CurrentRoute(request)

	urlParams := make(map[string]string) // creating url paramas map to pass it verify method in service to check validity of token
	urlParams["routeName"] = currentRoute.GetName()
	urlParams["eid"] = eid // storing manager id

	if authHeader != "" {
		urlParams["token"] = service.GetTokenFromHeader(authHeader) // Extracting the token from the header
		//fmt.Println(token)
		appErr := service.Verify(urlParams)
		if appErr != nil {
			writeResponse(response, http.StatusUnauthorized, notAuthorizedResponse(" Token Verification Failed "))
			fmt.Println(appErr)

		} else {

			result := service.CreateUserByID(userInfo)
			writeResponse(response, http.StatusOK, authorizedResponse())
			json.NewEncoder(response).Encode(result)
			fmt.Println("User Created")
			// call service to create user in database
		}

	} else {
		response.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(response, "missing token")
	}

}

func DeleteUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)
	eid := params["eid"] // this is the manager id who wants to create a employee under him which we need to verify with the token
	childEid := params["childEid"]

	authHeader := request.Header.Get("Authorization")
	currentRoute := mux.CurrentRoute(request)

	urlParams := make(map[string]string) // creating url paramas map to pass it verify method in service to check validity of token
	urlParams["routeName"] = currentRoute.GetName()
	urlParams["eid"] = eid // storing manager id

	if authHeader != "" {
		urlParams["token"] = service.GetTokenFromHeader(authHeader) // Extracting the token from the header
		//fmt.Println(token)
		appErr := service.Verify(urlParams)
		if appErr != nil {
			writeResponse(response, http.StatusUnauthorized, notAuthorizedResponse(" Token Verification Failed "))
			fmt.Println(appErr)

		} else {

			result := service.DeleteUserByID(childEid)
			writeResponse(response, http.StatusOK, authorizedResponse())
			json.NewEncoder(response).Encode(result)
			// call service to create user in database
		}

	} else {
		response.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(response, "missing token")
	}

}

func notAuthorizedResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"isAuthorized": false,
		"message":      msg,
	}
}

func authorizedResponse() map[string]bool {
	return map[string]bool{"isAuthorized": true}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
