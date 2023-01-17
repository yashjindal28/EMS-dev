package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func IsAuthorized(token string, routeName string, routeVars map[string]string) bool {

	url := buildVerifyURL(token, routeName, routeVars) // building url to call verify api of auth-service
	//fmt.Println(url)
	if response, err := http.Get(url); err != nil { // Calling get verify api call to auth service and storing its response
		fmt.Println("Error while verifying authrization..." + err.Error())
		return false // returning unauthorized
	} else {
		//fmt.Println(response)
		m := map[string]bool{}
		if err = json.NewDecoder(response.Body).Decode(&m); err != nil { // decoding response
			fmt.Println("Error while decoding response from auth service:" + err.Error())
			return false
		}
		return m["isAuthorized"] // returning the check of whter the token is authorized for the current route or not
	}

}

/*
  This will generate a url for token verification in the below format
  /auth/verify?token={token string}
              &routeName={current route name}
              &eid={Employee id from the current route}
              &desg={Designation from current route}
  Sample: /auth/verify?token=aaaa.bbbb.cccc&routeName=Payroll&eid=MCM002
*/

func buildVerifyURL(token string, routeName string, routeVars map[string]string) string {
	u := url.URL{Host: "localhost:8080", Path: "/auth/verify", Scheme: "http"} // creating api call path
	q := u.Query()
	q.Add("token", token)         // adding token to the query of api call
	q.Add("routeName", routeName) // adding current called route name to the api call query
	for k, v := range routeVars {
		q.Add(k, v) // adding vars of current route to query
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func GetTokenFromHeader(header string) string {
	/*
	   token is coming in the format as below
	   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50cyI6W.yI5NTQ3MCIsIjk1NDcyIiw"
	*/
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
