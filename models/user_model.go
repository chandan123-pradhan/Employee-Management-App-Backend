package models

import "go.mongodb.org/mongo-driver/bson/primitive"
type User struct{
	Name string `json:"name"`
	EmailId string `json:"email"`
	MobileNumber string `json:"mobile_number"`
	Description string `json:"description"`
	WalletAmount string `json:"wallet_amount"`
}


type UpdateWalletAmountModel struct{
	Id primitive.ObjectID `json:"_id"`
	WalletAmount string `json:"wallet_amount"`
	
}