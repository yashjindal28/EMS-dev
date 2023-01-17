package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type PayrollData struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	EmployeeID    string             `json:"eid,omitempty" bson:"eid,omitempty"`
	Firstname     string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname      string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Designation   string             `json:"desg,omitempty" bson:"desg,omitempty"`
	DepartmentID  string             `json:"did,omitempty" bson:"did,omitempty"`
	Department    string             `json:"dpt,omitempty" bson:"dpt,omitempty"`
	Manager       string             `json:"manager,omitempty" bson:"manager,omitempty"`
	ManagerID     string             `json:"managerID,omitempty" bson:"managerID,omitempty"`
	Salary        string             `json:"salary,omitempty" bson:"salary,omitempty"`
	Email         string             `json:"email,omitempty" bson:"email,omitempty"`
	IsCheckIssued int                `json:"isCheckIssued" bson:"isCheckIssued"`
}
