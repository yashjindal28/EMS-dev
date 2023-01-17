package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	EmployeeID   string             `json:"eid,omitempty" bson:"eid,omitempty"`
	Firstname    string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname     string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Gender       string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Designation  string             `json:"desg,omitempty" bson:"desg,omitempty"`
	DepartmentID string             `json:"did,omitempty" bson:"did,omitempty"`
	Department   string             `json:"dpt,omitempty" bson:"dpt,omitempty"`
	Manager      string             `json:"manager,omitempty" bson:"manager,omitempty"`
	ManagerID    string             `json:"managerID,omitempty" bson:"managerID,omitempty"`
	Salary       string             `json:"salary,omitempty" bson:"salary,omitempty"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty"`
	Location     string             `json:"location,omitempty" bson:"location,omitempty"`
}

type PersonalInfo struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	EmployeeID string             `json:"eid,omitempty" bson:"eid,omitempty"`
	Address    string             `json:"address,omitempty" bson:"address,omitempty"`
	Contact    string             `json:"contact,omitempty" bson:"contact,omitempty"`
}

type EmployeeLeave struct {
	EmployeeID      string `json:"eid,omitempty" bson:"eid,omitempty"`
	Firstname       string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname        string `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Manager         string `json:"manager,omitempty" bson:"manager,omitempty"`
	ManagerID       string `json:"managerID,omitempty" bson:"managerID,omitempty"`
	LeavesRemaining int    `json:"leavesRem,omitempty" bson:"leavesRem,omitempty"`
	DaysRequested   int    `json:"daysReq" bson:"daysReq"`
	PendingStatus   int    `json:"pendingStatus" bson:"pendingStatus"`
	Approved        int    `json:"approved" bson:"approved"`
}

type Payroll struct {
	EmployeeID    string `json:"eid,omitempty" bson:"eid,omitempty"`
	IsCheckIssued int    `json:"isCheckIssued" bson:"isCheckIssued"`
}

type Count struct {
	CountOfEmp int `json:"countOfEmp" bson:"countOfEmp"`
	CountOfDpt int `json:"countOfDpt" bson:"countOfDpt"`
}

//db.count.insert({"countOfDpt":1,"countOfEmp":1})

type Department struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DepartmentID   string             `json:"did,omitempty" bson:"did,omitempty"`
	DepartmentName string             `json:"dpt,omitempty" bson:"dpt,omitempty"`
	Manager        string             `json:"manager,omitempty" bson:"manager,omitempty"`
	ManagerID      string             `json:"managerID,omitempty" bson:"managerID,omitempty"`
	EmployeesOfDpt []EmployeesOfDpt   `json:"eiddpt,omitempty" bson:"eiddpt,omitempty"`
}

type EmployeesOfDpt struct {
	EmployeeID        string `json:"eid,omitempty" bson:"eid,omitempty"`
	AssignedToProject string `json:"assignedTo,omitempty" bson:"assignedTo,omitempty"`
}
