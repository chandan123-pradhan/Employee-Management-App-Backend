package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chandan123-pradhan/employee_management_backend/routers"
)


func main()  {
	r := routers.Routers()

	log.Fatal(http.ListenAndServe(":8080",r))
	fmt.Println("Listening port is 8080")
}