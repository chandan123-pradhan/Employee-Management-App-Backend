package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/chandan123-pradhan/employee_management_backend/helpers"
	"github.com/chandan123-pradhan/employee_management_backend/models"
)

func AddEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var employee models.Employee

	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		helpers.OperationSuccess(w, "You are sending data in wrong format kindly please Recheck your request json")
	}
	if employee.Name == "" || employee.Name == " " {
		helpers.OperationSuccess(w, "Name is Required")
	} else if employee.EmailId == "" || employee.EmailId == " " {
		helpers.OperationSuccess(w, "Email id Required")
	} else if employee.MobileNumber == "" || employee.MobileNumber == " " {
		helpers.OperationSuccess(w, "Mobile Number is Required")
	} else {

		var result bool= addEmployeesInDb(employee)
		if(result){
			helpers.OperationSuccess(w, "Employees Added Successfully Done")

		}else{
			helpers.OperationSuccess(w, "Operation Faild please try again")
		}
		
	}
}


func GetAllEmployees(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	result := getAllEmployeesFromDb()
	helpers.OperationSuccessWithData(w,result)
		
}
