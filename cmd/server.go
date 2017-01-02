package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/u8slvn/raspigosms/raspigosms"
)

var router *mux.Router

// StartServer builds the router and then start to listen on the
// given http address.
func StartServer() {
	fmt.Println("Server starting...")

	smsHandler := newSmsHandler()

	router = mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/sms", smsHandler.create).Methods("POST")
	api.HandleFunc("/sms", smsHandler.index).Methods("GET")
	api.HandleFunc("/sms/{id}", smsHandler.show).Methods("GET").Name("sms_show")

	srv := &http.Server{
		Handler:      router,
		Addr:         raspigosms.Conf.HTTPAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
