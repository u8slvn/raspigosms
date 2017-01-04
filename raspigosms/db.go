package raspigosms

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

func dbConnect() {
	conn, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	db = conn.DB("raspi_go_sms")
	fmt.Println("Database connected.")
}
