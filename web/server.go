package web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var router *mux.Router

// StartServer builds the router and then start to listen on the
// given http address.
func StartServer(HTTPAddr string) {
	fmt.Println("Server starting...")

	smsController := NewSmsController()

	router = mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/sms", smsController.Create).Methods("POST")
	api.HandleFunc("/sms", smsController.Index).Methods("GET")
	api.HandleFunc("/sms/{id}", smsController.Show).Methods("GET").Name("sms_show")

	srv := &http.Server{
		Handler:      router,
		Addr:         HTTPAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
