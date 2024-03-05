package routers

import (
	"github.com/chandan123-pradhan/employee_management_backend/controllers"
	"github.com/gorilla/mux"
)


func Routers() *mux.Router  {
	router := mux.NewRouter()

	router.HandleFunc("/api/add_employee",controllers.AddEmployees)
	router.HandleFunc("/api/get_employees",controllers.GetAllEmployees)
	router.HandleFunc("/api/delete_employee/{id}",controllers.DeleteEmployee)
	router.HandleFunc("/api/update_amount",controllers.UpdateUsersWalletAmount).Methods("POST")
	return router
}