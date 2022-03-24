package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/*MongoCN is the connection object to the database*/
var MongoCN = ConnectDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://crision98:ionutu89@lisander.8nory.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

//var clientOptions = options.Client().ApplyURI("mongodb://admin:admin@192.168.1.138:27017/?connect=direct")

/*ConnectDB is the function that allows me to connect to the database.*/
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
		return client
	}
	log.Println("Successful connection with the database")
	return client
}

/*CheckConnection is the Ping to the database*/
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return 0
	}
	return 1
}
