package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Departments struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DepartmentID   string             `json:"did,omitempty" bson:"did,omitempty"`
	DepartmentName string             `json:"dpt,omitempty" bson:"dpt,omitempty"`
	Manager        string             `json:"manager,omitempty" bson:"manager,omitempty"`
	ManagerID      string             `json:"managerID,omitempty" bson:"managerID,omitempty"`
	EmployeesOfDpt []EmployeesOfDpt   `json:"eiddpt,omitempty" bson:"eiddpt,omitempty"`
	Projects       []Projects         `json:"projects,omitempty" bson:"projects,omitempty"`
}

type Projects struct {
	ProjectID        string             `json:"pid,omitempty" bson:"pid,omitempty"`
	ProjectName      string             `json:"projectName,omitempty" bson:"projectName,omitempty"`
	Deadline         string             `json:"deadline,omitempty" bson:"deadline,omitempty"`
	ProjectEmployees []ProjectEmployees `json:"projectEmployees" bson:"projectEmployees"`
}

type ProjectEmployees struct {
	EmployeeID string `json:"eid" bson:"eid"`
}

type EmployeesOfDpt struct {
	EmployeeID        string `json:"eid,omitempty" bson:"eid,omitempty"`
	AssignedToProject string `json:"assignedTo,omitempty" bson:"assignedTo,omitempty"`
	Firstname         string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname          string `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Email             string `json:"email,omitempty" bson:"email,omitempty"`
	Designation       string `json:"desg,omitempty" bson:"desg,omitempty"`
}

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

type CustomObj struct {
	EmployeeID   string `json:"eid,omitempty" bson:"eid,omitempty"`
	DepartmentID string `json:"did,omitempty" bson:"did,omitempty"`
	ProjectID    string `json:"pid,omitempty" bson:"pid,omitempty"`
}

type Count struct {
	CountOfEmp int `json:"countOfEmp" bson:"countOfEmp"`
	CountOfDpt int `json:"countOfDpt" bson:"countOfDpt"`
}

//db.count.insert({"countOfDpt":1,"countOfEmp":1})
