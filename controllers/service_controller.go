package controllers

import (
	"context"
	"fmt"
	"github.com/chandan123-pradhan/employee_management_backend/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const connnectionString = "mongodb+srv://chandanpradhanbxr:I0rBik4Fb86FHjOM@cluster0.xbvad91.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0&ipv6=false" //mongo db connection url
const dbName = "Employee Management"
const colName = "Employees List"
var collection *mongo.Collection

func init()  {
	clientOption := options.Client().ApplyURI(connnectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	helpers.CheckIfNull(err)
	fmt.Println("Mongodb connection success")
	collection=client.Database(dbName).Collection(colName)
	fmt.Println("Connections instant is ready")
}






