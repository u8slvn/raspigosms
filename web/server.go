package web

import (
	"log"
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
	"github.com/u8slvn/raspigosms/web/controllers"
)

// StartServer builds the router and then start to listen on the
// given http address
func StartServer(HTTPAddr string) {

	//Init controllers
	smsController := controllers.NewSmsController(getSession())

	// Routing
	router := mux.NewRouter()

	// api Subrouter
	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/sms", smsController.Create).Methods("POST")
	api.HandleFunc("/sms", smsController.Index).Methods("GET")

	// Start the web server
	srv := &http.Server{
		Handler: router,
		Addr:    HTTPAddr,
		// Enforce timeouts for servers
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func getSession() *mgo.Session {
	// Connect to local mongo
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	return session
}
