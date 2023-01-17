package service

import (
	"context"
	"fmt"
	"net/smtp"
	"strconv"
	"time"

	"github.com/yashjindal28/go-EMS-payrollService/model"
	"github.com/yashjindal28/go-EMS-payrollService/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAll() (err error, err2 error, payrollData []model.PayrollData) {
	// db.payrollInfo.updateOne({"eid":"005MCMP"},{$set:{"isCheckIssued":0}})
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//cursor, err := repository.FindAllPayrollData()
	cursor, err := repository.FindAllPayrollData()
	for cursor.Next(ctx) {
		var payInfo model.PayrollData
		cursor.Decode(&payInfo)
		payrollData = append(payrollData, payInfo)
	}
	err2 = cursor.Err()
	if err2 != nil {
		panic(err2)
	}

	return err, err2, payrollData

}

func AddPayrollInfo(employee model.PayrollData) *mongo.InsertOneResult {
	result := repository.AddPayrollInfo(employee)

	return result
}

func DeleteEmployeeByID(childEid string) *mongo.DeleteResult {
	result := repository.DeleteEmployeeByID(childEid)

	return result
}

func CalSalaryAndMail(eid string) {
	var sal, email string
	var payInfo model.PayrollData

	repository.GetPayrollInfoOfEmployeeByID(eid).Decode(&payInfo)
	repository.UpdatePayrollInfo(eid)

	sal = payInfo.Salary
	email = payInfo.Email

	salary, _ := strconv.ParseFloat(sal, 64)
	salary = salary - (salary * 0.10)
	salary = (salary / 12)
	//fmt.Println(salary)
	mail(email, salary)
}

func mail(email string, salary float64) {
	// Sender data.
	from := "alertnotificationB14@gmail.com"
	password := "kletseuhftcpasmv"
	//
	//fmt.Println("Inside")
	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

func UpdateEmployeeByIdUnderManager(eid string, employee model.PayrollData) (result *mongo.UpdateResult, err error) {

	result, err = repository.UpdateEmployeeByIdUnderManager(eid, employee)

	return result, err
}
