package web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/u8slvn/raspigosms/web/controllers"
)

// StartServer builds the router and then start to listen on the
// given http address.
func StartServer(HTTPAddr string) {
	fmt.Println("Server starting...")

	smsController := controllers.NewSmsController()

	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/sms", smsController.Create).Methods("POST")
	api.HandleFunc("/sms", smsController.Index).Methods("GET")
	api.HandleFunc("/sms/{id}", smsController.Show).Methods("GET")

	srv := &http.Server{
		Handler:      router,
		Addr:         HTTPAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
