package repository

import (
	"context"
	"log"
	"time"

	"github.com/yashjindal28/go-EMS-departmentService/database"
	"github.com/yashjindal28/go-EMS-departmentService/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client = database.ConnectDB()

func GetAllDepartmentData() (cursor *mongo.Cursor, err error) {

	collection := client.Database("department_db").Collection("departments")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err = collection.Find(ctx, bson.M{})

	return cursor, err
}

func AddNewEmployeeToDepartment(employee model.Employee) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var count model.Count
	collection4 := client.Database("department_db").Collection("count")
	err := collection4.FindOne(ctx, bson.M{}).Decode(&count)
	if err != nil {
		panic(err)
	}

	if employee.ManagerID[:3] == "EXE" {
		//create a new department

		var dpt model.Departments
		//dpt.DepartmentID = "D" + strconv.Itoa(count.CountOfDpt+1)
		dpt.DepartmentID = employee.DepartmentID
		dpt.DepartmentName = employee.Department
		dpt.Manager = employee.Firstname + " " + employee.Lastname
		dpt.ManagerID = employee.EmployeeID
		collection3 := client.Database("department_db").Collection("departments")
		collection3.InsertOne(ctx, dpt)

		collection3 = client.Database("department_db").Collection("count")
		filter := bson.M{}
		update := bson.M{
			"$set": bson.M{
				"countOfDpt": count.CountOfDpt + 1,
				"countOfEmp": count.CountOfEmp + 1,
			},
		}
		_, err = collection3.UpdateOne(ctx, filter, update) // you can simply replace using replace one command and decoded personalInfo obejct
		if err != nil {
			panic(err)
		}

	} else if employee.ManagerID == "admin" {
		// adding CEO of organisation
		var dpt model.Departments
		dpt.DepartmentID = "D1"
		employee.DepartmentID = dpt.DepartmentID
		dpt.DepartmentName = employee.Department
		dpt.Manager = employee.Firstname + " " + employee.Lastname
		dpt.ManagerID = employee.EmployeeID

		collection3 := client.Database("department_db").Collection("departments")
		collection3.InsertOne(ctx, dpt)

		collection3 = client.Database("department_db").Collection("count")
		filter := bson.M{}
		update := bson.M{
			"$set": bson.M{
				"countOfDpt": 1,
				"countOfEmp": 1,
			},
		}
		_, err = collection3.UpdateOne(ctx, filter, update) // you can simply replace using replace one command and decoded personalInfo obejct
		if err != nil {
			panic(err)
		}

	} else {
		// add a employee to a existing department

		//var department Department
		collection3 := client.Database("department_db").Collection("departments")

		change := bson.M{"$push": bson.M{"eiddpt": bson.M{"eid": employee.EmployeeID, "assignedTo": "none", "firstname": employee.Firstname, "lastname": employee.Lastname, "email": employee.Email, "desg": employee.Designation}}}
		match := bson.M{"did": employee.DepartmentID}
		_, err = collection3.UpdateOne(ctx, match, change)

		if err != nil {
			panic(err)
		}

		collection3 = client.Database("department_db").Collection("count")
		filter := bson.M{}
		update := bson.M{
			"$set": bson.M{
				"countOfEmp": count.CountOfEmp + 1,
			},
		}
		_, err = collection3.UpdateOne(ctx, filter, update) // you can simply replace using replace one command and decoded personalInfo obejct
		if err != nil {
			panic(err)
		}
	}
}

func DeleteEmployeeByID(childEid string, did string) *mongo.UpdateResult {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("department_db").Collection("departments")

	var department model.Departments
	err := collection.FindOne(ctx, bson.M{"did": did, "eiddpt.eid": childEid}).Decode(&department)

	var pid string
	for _, value := range department.EmployeesOfDpt {
		if value.EmployeeID == childEid {
			pid = value.AssignedToProject
		}
	}

	selector := bson.M{"did": did}
	update := bson.M{"$pull": bson.M{"eiddpt": bson.M{"eid": childEid}}}
	result, err := collection.UpdateOne(ctx, selector, update)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(pid)
	selector = bson.M{"did": did, "projects.pid": pid}
	update = bson.M{"$pull": bson.M{"projects.$.projectEmployees": bson.M{"eid": childEid}}}
	result, err = collection.UpdateOne(ctx, selector, update)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(result)
	return result
}

func GetDepartmentByID(did string) (department *mongo.SingleResult) {

	collection := client.Database("department_db").Collection("departments")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	department = collection.FindOne(ctx, bson.M{"did": did})
	return department
}

func SaveProjectInDepartment(did string, project model.Projects) {

	collection := client.Database("department_db").Collection("departments")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	change := bson.M{"$push": bson.M{"projects": bson.M{"pid": project.ProjectID, "deadline": project.Deadline, "projectName": project.ProjectName}}}
	match := bson.M{"did": did}
	_, err := collection.UpdateOne(ctx, match, change)

	if err != nil {
		panic(err)
	}
}

func GetEmployeeDetailsUnderDepartment(did string) (department model.Departments, err error) {

	//var department model.Departments
	collection := client.Database("department_db").Collection("departments")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = collection.FindOne(ctx, bson.M{"did": did}).Decode(&department)

	//var eids []string
	//for _, value := range department.EmployeesOfDpt {
	//	eids = append(eids, value.EmployeeID)
	//}

	//collection = client.Database("employees_db").Collection("employees")
	//cursor, err := collection.Find(ctx, bson.M{"eid": bson.M{"$in": eids}})
	return department, err
}

func SaveEmployeeToAProject(customObj model.CustomObj) {

	collection := client.Database("department_db").Collection("departments")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//_ = DeleteEmployeeByID(customObj.EmployeeID, customObj.DepartmentID)
	var department model.Departments
	err := collection.FindOne(ctx, bson.M{"did": customObj.DepartmentID, "eiddpt.eid": customObj.EmployeeID}).Decode(&department)

	var pid string
	for _, value := range department.EmployeesOfDpt {
		if value.EmployeeID == customObj.EmployeeID {
			pid = value.AssignedToProject
		}
	}

	selector := bson.M{"did": customObj.DepartmentID, "projects.pid": pid}
	update := bson.M{"$pull": bson.M{"projects.$.projectEmployees": bson.M{"eid": customObj.EmployeeID}}}
	_, err = collection.UpdateOne(ctx, selector, update)
	if err != nil {
		log.Fatal(err)
	}

	// Deleted previous entries of employee , now updating new ones

	change := bson.M{"$push": bson.M{"projects.$.projectEmployees": bson.M{"eid": customObj.EmployeeID}}}
	match := bson.M{"did": customObj.DepartmentID, "projects.pid": customObj.ProjectID}
	_, err = collection.UpdateOne(ctx, match, change)

	if err != nil {
		panic(err)
	}
	change = bson.M{"$set": bson.M{"eiddpt.$.assignedTo": customObj.ProjectID}}
	match = bson.M{"did": customObj.DepartmentID, "eiddpt.eid": customObj.EmployeeID}
	_, err = collection.UpdateOne(ctx, match, change)

	if err != nil {
		panic(err)
	}
}

func GetCount() model.Count {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var count model.Count
	collection4 := client.Database("department_db").Collection("count")
	err := collection4.FindOne(ctx, bson.M{}).Decode(&count)
	if err != nil {
		panic(err)
	}

	return count
}

func UpdateEmployeeByIdUnderManager(eid string, employee model.Employee) (result *mongo.UpdateResult, err error) {

	collection := client.Database("department_db").Collection("departments")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	change := bson.M{"$set": bson.M{"eiddpt.$.firstname": employee.Firstname, "eiddpt.$.lastname": employee.Lastname, "eiddpt.$.email": employee.Email, "eiddpt.$.desg": employee.Designation}}
	//change := bson.M{"$set": bson.M{"eiddpt": bson.M{"firstname": employee.Firstname, "lastname": employee.Lastname, "email": employee.Email, "desg": employee.Designation}}}
	match := bson.M{"did": employee.DepartmentID, "eiddpt.eid": employee.EmployeeID}
	_, err = collection.UpdateOne(ctx, match, change)

	if err != nil {
		panic(err)
	}

	filter := bson.D{{"managerID", eid}}
	update := bson.M{
		"$set": bson.M{
			"manager": employee.Firstname + " " + employee.Lastname,
		},
	}
	_, err = collection.UpdateOne(ctx, filter, update) // you can simply replace using replace one command and decoded personalInfo obejct
	if err != nil {
		panic(err)
	}

	return result, err
}
