package model

type EmployeeLeave struct {
	EmployeeID      string `json:"eid,omitempty" bson:"eid,omitempty"`
	Firstname       string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname        string `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Manager         string `json:"manager,omitempty" bson:"manager,omitempty"`
	ManagerID       string `json:"managerID,omitempty" bson:"managerID,omitempty"`
	LeavesRemaining int    `json:"leavesRem,omitempty" bson:"leavesRem,omitempty"`
	StartingDate    string `json:"sdate,omitempty" bson:"sdate,omitempty"`
	EndingDate      string `json:"edate,omitempty" bson:"edate,omitempty"`
	DaysRequested   int    `json:"daysReq" bson:"daysReq"`
	Reason          string `json:"reason" bson:"reason"`
	PendingStatus   int    `json:"pendingStatus" bson:"pendingStatus"`
	Approved        int    `json:"approved" bson:"approved"`
}
