package database

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

// DBConnection is the main connection handle for the database.
var DBConnection *mgo.Database

// Connect to local mongo.
func Connect() {
	Connection, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	DBConnection = Connection.DB("raspi_go_sms")
	fmt.Println("Database connected.")
}
