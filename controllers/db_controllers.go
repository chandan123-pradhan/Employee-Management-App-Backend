package controllers

import (
	"context"
	"fmt"

	"github.com/chandan123-pradhan/employee_management_backend/helpers"
	"github.com/chandan123-pradhan/employee_management_backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connnectionString = "mongodb+srv://chandanpradhanbxr:I0rBik4Fb86FHjOM@cluster0.xbvad91.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0&ipv6=false" //mongo db connection url
const dbName = "EmployeeManagement"
const colName = "Employees List"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connnectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	helpers.CheckIfNull(err)
	fmt.Println("Mongodb connection success")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Connections instant is ready")
}

func addEmployeesInDb(employee models.User) bool {
	res, err := collection.InsertOne(context.Background(), employee)
	fmt.Println(res.InsertedID);
	return helpers.CheckIfNull(err)

}

func getAllEmployeesFromDb() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.M{})
	helpers.CheckIfNull(err)
	var employees []primitive.M

	for cur.Next(context.Background()) {
		var employee bson.M
		err := cur.Decode(&employee)
		helpers.CheckIfNull(err)
		employees = append(employees, employee)
	}
	defer cur.Close(context.Background())
	return employees
}

func deleteEmployee(employeeId string) {
	id, _ := primitive.ObjectIDFromHex(employeeId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	helpers.CheckIfNull(err)
	fmt.Println("Employee deleted, with count =", deleteCount)
}

func updateUserWalletAmount(userId string, currentAmount string) bool {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		fmt.Println("user id not found")
		return false
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"walletamount": currentAmount}}
	_, e := collection.UpdateOne(context.Background(), filter, update)
	return helpers.CheckIfNull(e)

}

func getUserDetails(userId string) models.User {
	id, err := primitive.ObjectIDFromHex(userId)
	helpers.CheckIfNull(err)
	filter := bson.M{"_id": id}

	var user models.User
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Handle case where user with given userID is not found
			helpers.CheckIfNull(err)
		}

	}

	return user
}
