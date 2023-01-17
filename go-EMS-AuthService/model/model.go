package model

type LoginRequest struct {
	EmployeeID string `json:"eid,omitempty" bson:"eid,omitempty"`
	Password   string `json:"password" bson:"password"`
}

type Login struct {
	EmployeeID string   `json:"eid,omitempty" bson:"eid,omitempty"`
	Employee   Employee `json:"info,omitempty" bson:"info,omitempty"`
}

type Employee struct {
	Designation string `json:"desg,omitempty" bson:"desg,omitempty"`
}

type Claims struct {
	EmployeeID  string `json:"eid,omitempty" bson:"eid,omitempty"`
	Designation string `json:"desg,omitempty" bson:"desg,omitempty"`
	Expiry      int64  `json:"exp"`
}

type RolePermissions struct {
	rolePermissions map[string][]string
}

type UserInfo struct {
	EmployeeID  string `json:"eid,omitempty" bson:"eid,omitempty"`
	Designation string `json:"desg,omitempty" bson:"desg,omitempty"`
}
