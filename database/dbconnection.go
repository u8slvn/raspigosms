package database

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

// DBConnection is the main connection handle for the database.
var DBConnection *mgo.Session

// Connect to local mongo.
func Connect() {
	var err error
	DBConnection, err = mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected.")
}
