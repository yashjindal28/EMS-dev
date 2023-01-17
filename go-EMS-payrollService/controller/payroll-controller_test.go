package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/payroll/{eid}", AllEmployeeEndpoint).Methods("GET").Name("GetPayrollDataForAllEmployess")
	router.HandleFunc("/payroll/{eid}/add-info", AddPayrollInfoEndpoint).Methods("POST").Name("AddPayrollInfo")
	router.HandleFunc("/payroll/{eid}/delete-info/{childEid}", DeletePayrollInfoForEmployeeEndpoint).Methods("DELETE").Name("DeleteEmployee")
	router.HandleFunc("/payroll/{eid}/issue/{childEid}", IssuePaycheckEndpoint).Methods("GET").Name("IssuePayCheckByEmployeeID")

	return router
}

func TestAddPayrollInfoEndpoint(t *testing.T) {
	data := url.Values{}
	data.Set("eid", "test")
	data.Set("firstname", "test")
	data.Set("salary", "10000")
	data.Set("email", "yashjindal50000@gmail.com")
	request, _ := http.NewRequest("POST", "/payroll/admin/add-info", strings.NewReader(data.Encode()))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
