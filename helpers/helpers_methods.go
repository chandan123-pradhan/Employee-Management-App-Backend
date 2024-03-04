package helpers

import "log"


func CheckIfNull(err error )  {
	if(err!=nil){
		log.Fatal(err)
	}
}