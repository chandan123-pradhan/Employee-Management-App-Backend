package helpers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/chandan123-pradhan/employee_management_backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckIfNull(err error)bool {
	if err != nil {
		log.Fatal(err)
		return false
	}else {
		return true
	}
}

func OperationSuccess(w http.ResponseWriter,message string) {
	var baseResponse = models.BaseResponse{
		StatusCode: 200,
		Message:    message,
		Data:       map[string]interface{}{},
	}

	json.NewEncoder(w).Encode(baseResponse)
}


func UnAutherizedError(w http.ResponseWriter) {
	var baseResponse = models.BaseResponse{
		StatusCode: 400,
		Message:    "You are Unautherized kindly please login first",
		Data:       map[string]interface{}{},
	}

	json.NewEncoder(w).Encode(baseResponse)
}


func OperationSuccessWithData(w http.ResponseWriter, data []primitive.M ) {
	var baseResponse = models.BaseResponse{
		StatusCode: 200,
		Message:    "Employees List Fetchded successfully done",
		Data:       data,
	}

	json.NewEncoder(w).Encode(baseResponse)
}
