package main

import (
	"github.com/u8slvn/raspigosms/app"
	"github.com/u8slvn/raspigosms/database"
	"github.com/u8slvn/raspigosms/web"
)

func main() {
	database.Connect()
	app.StartWorking()
	web.StartServer()
}
