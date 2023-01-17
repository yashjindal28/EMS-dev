package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/yashjindal28/go-EMS-AuthService/model"
	"github.com/yashjindal28/go-EMS-AuthService/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

const TOKEN_DURATION = time.Hour
const HMAC_SAMPLE_SECRET = "abcdefgh"

func Login(loginRequest model.LoginRequest) (*string, string, error) {
	// Step 1 - Verify the credentials i.e username and password
	// Step 2 - generate the token

	login, err := repository.FindBy(loginRequest)
	if err != nil { // if creds are invalid return unauth error
		return nil, "", err
	}

	token, err := GenerateToken(login)

	return token, login.Employee.Designation, err
}

func GenerateToken(login *model.Login) (*string, error) {
	var claims = jwt.MapClaims{
		"eid":  login.EmployeeID,
		"desg": login.Employee.Designation,
		"exp":  time.Now().Add(TOKEN_DURATION).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenAsString, err := token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		log.Println("Failed while generating token" + err.Error())
		return nil, errors.New("cannot generate token")
	}

	return &signedTokenAsString, nil
}

// 3 Steps --
// 1) Validity of the token (includes expiry and signature verification)
// 2) Verify if the role has access to the resource
// 3) need to verify if the resource being accessd if for the same user

func Verify(urlParams map[string]string) error {

	// Convert the string token to struct token to check validation
	jwtToken, err := jwtTokenFromString(urlParams["token"])
	//fmt.Println(jwtToken)
	if err != nil {
		return errors.New("Not able to convert string token to jwt struct" + err.Error())
	} else {
		/*
		   Step 1) Checking the validity of the token, jwtToken.Valid verifies the expiry
		   time and the signature of the token
		*/
		if jwtToken.Valid {
			// type cast the token claims to jwt.MapClaims
			mapClaims := jwtToken.Claims.(jwt.MapClaims)
			// converting token claims to claims struct that we can use
			claims, err := buildStructClaimFromMapClaims(mapClaims)
			if err != nil {
				return errors.New("Not able to convert mapClaims to claims struct" + err.Error())
			}

			// Step 3) Check if eid in token matches to eid is route parameter
			if claims.EmployeeID != urlParams["eid"] {
				return errors.New("EID doesn't match in route parameter and token - Unauthorized Access")
			}

			// Step 2) Checking the role based access
			isAuthorized := isAuthorizedFor(claims.Designation, urlParams["routeName"])
			if !isAuthorized {
				return errors.New(fmt.Sprintf("%s role is not authorized", claims.Designation))
			}

			return nil // returning nil if all conditions satisfy
		} else {
			return errors.New("invalid Token - Unauthorized Access")
		}
	}
}

func jwtTokenFromString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		log.Println("Error while parsing token: " + err.Error())
		return nil, err
	}
	return token, nil
}

func buildStructClaimFromMapClaims(mapClaims jwt.MapClaims) (*model.Claims, error) {
	bytes, err := json.Marshal(mapClaims)
	if err != nil {
		return nil, err
	}
	var claims model.Claims
	err = json.Unmarshal(bytes, &claims)
	if err != nil {
		return nil, err
	}
	return &claims, nil
}

func isAuthorizedFor(desg string, routeName string) bool {

	rolePermissions := map[string][]string{
		"Manager": {"AddPayrollInfo", "GetPayrollDataForAllEmployess", "IssuePayCheckByEmployeeID", "GetCount", "AddNewEmployeeToDepartment", "GetAllDepartmentsData", "GetADepartmentByID", "GetEmployeeDetailsOfADepartment", "SaveProjectInADepartment", "SaveEmployeeToAProject", "AddLeaveDataForNewEmployee", "EmployeeLeaveData", "EmployeeRequestLeave", "AllEmployeesUnderManagerLeaveData", "AprroveLeave", "RejectLeave", "CreateEmployee", "GetAllEmployeesUnderManager", "GetOneEmployeeDetails", "DeleteEmployee", "UpdateEmployee", "GetEmployeePersonalDataByID", "UpdatePersonalInfo", "AllEmployeesInCompany", "Search", "Filter", "CreateUser", "DeleteUser"},
		"CEO":     {"AddPayrollInfo", "GetPayrollDataForAllEmployess", "IssuePayCheckByEmployeeID", "GetCount", "AddNewEmployeeToDepartment", "GetAllDepartmentsData", "GetADepartmentByID", "GetEmployeeDetailsOfADepartment", "SaveProjectInADepartment", "SaveEmployeeToAProject", "AddLeaveDataForNewEmployee", "AllEmployeesUnderManagerLeaveData", "AprroveLeave", "RejectLeave", "CreateEmployee", "GetAllEmployeesUnderManager", "GetOneEmployeeDetails", "DeleteEmployee", "UpdateEmployee", "GetEmployeePersonalDataByID", "UpdatePersonalInfo", "AllEmployeesInCompany", "Search", "Filter", "CreateUser", "DeleteUser"},
		"admin":   {"AddPayrollInfo", "GetPayrollDataForAllEmployess", "IssuePayCheckByEmployeeID", "GetCount", "AddNewEmployeeToDepartment", "GetAllDepartmentsData", "GetADepartmentByID", "GetEmployeeDetailsOfADepartment", "SaveProjectInADepartment", "SaveEmployeeToAProject", "AddLeaveDataForNewEmployee", "AllEmployeesUnderManagerLeaveData", "AprroveLeave", "RejectLeave", "CreateEmployee", "GetAllEmployeesUnderManager", "GetOneEmployeeDetails", "DeleteEmployee", "UpdateEmployee", "GetEmployeePersonalDataByID", "UpdatePersonalInfo", "AllEmployeesInCompany", "Search", "Filter", "CreateUser", "DeleteUser"},
		"other":   {"EmployeeLeaveData", "EmployeeRequestLeave", "GetEmployeePersonalDataByID", "UpdatePersonalInfo", "AllEmployeesInCompany", "Search", "Filter", "GetOneEmployeeDetails"},
	}
	if desg != "Manager" && desg != "CEO" && desg != "admin" {
		desg = "other"
	}
	perms := rolePermissions[desg]
	for _, r := range perms {
		if r == strings.TrimSpace(routeName) {
			return true
		}
	}
	return false
}

func CreateUserByID(userInfo model.UserInfo) *mongo.InsertOneResult {
	result := repository.CreateUserByID(userInfo)

	return result
}

func DeleteUserByID(childEid string) *mongo.DeleteResult {
	result := repository.DeleteUserByID(childEid)

	return result
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
