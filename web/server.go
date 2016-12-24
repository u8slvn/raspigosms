package web

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// StartServer builds the router and then start to listen on the
// given http address
func StartServer(HTTPAddr string) {
	r := mux.NewRouter()

	// api Subrouter
	apir := r.PathPrefix("/api").Subrouter()
	apir.HandleFunc("/sms", Collector).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    HTTPAddr,
		// Enforce timeouts for servers
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
