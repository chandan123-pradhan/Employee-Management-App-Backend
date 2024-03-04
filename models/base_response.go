package models

type BaseResponse struct{
	StatusCode int         `json:"statusCode"`
    Message    string      `json:"message"`
    Data       interface{} `json:"data"`
}