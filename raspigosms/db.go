package raspigosms

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

// DBConnection is the main connection handle for the database.
var db *mgo.Database

func dbConnect() {
	connection, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	db = connection.DB("raspi_go_sms")
	fmt.Println("Database connected.")
}
