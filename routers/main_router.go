package routers

import (
	"github.com/chandan123-pradhan/employee_management_backend/controllers"
	"github.com/gorilla/mux"
)


func Routers() *mux.Router  {
	router := mux.NewRouter()

	router.HandleFunc("/api/add_employee",controllers.AddEmployees)
	router.HandleFunc("/api/get_employees",controllers.GetAllEmployees)
	return router
}