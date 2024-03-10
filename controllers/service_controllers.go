package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/chandan123-pradhan/employee_management_backend/helpers"
	"github.com/chandan123-pradhan/employee_management_backend/models"
	"github.com/gorilla/mux"
)

func AddEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var employee models.User

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
	} else if employee.Description == "" || employee.Description == " " {
		helpers.OperationSuccess(w, "Description is Required")
	} else {

		var result bool = addEmployeesInDb(employee)
		if result {
			helpers.OperationSuccess(w, "Employees Added Successfully Done")

		} else {
			helpers.OperationSuccess(w, "Operation Faild please try again")
		}

	}
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	result := getAllEmployeesFromDb()
	helpers.OperationSuccessWithData(w, result)

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteEmployee(params["id"])
	helpers.OperationSuccess(w, "Employee Deleted Successfully Done")

}

func UpdateUsersWalletAmount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var employee models.UpdateWalletAmountModel

	err := json.NewDecoder(r.Body).Decode(&employee)

	helpers.CheckIfNull(err)
	var result bool =updateUserWalletAmount(employee.Id.Hex(),employee.WalletAmount)
	if(result){
		helpers.OperationSuccess(w,"Wallet amount updated sucecssfully done")
	}else{
		helpers.OperationSuccess(w,"Wallet Amount update request faild")
	}

}


func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	params := mux.Vars(r)
	result :=getUserDetails(params["id"])

	helpers.OperationSuccessWithObject(w, result)

}