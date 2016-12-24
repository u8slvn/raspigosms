package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nu7hatch/gouuid"
	"github.com/u8slvn/raspigosms/app"
	"github.com/u8slvn/raspigosms/database"
	"github.com/u8slvn/raspigosms/gsm"
)

// SmsController provides the sms routes handlers.
type SmsController struct{}

// NewSmsController creates and return a new SmsController.
func NewSmsController() *SmsController {
	return &SmsController{}
}

// Create collect post data to create Sms wich sent onto the SmsQueue channel.
func (sc *SmsController) Create(w http.ResponseWriter, r *http.Request) {
	// Retrieve the sms information from the request.
	phone := r.FormValue("phone")
	message := r.FormValue("message")

	// Try to create the sms.
	sms, err := gsm.NewSms(phone, message, gsm.SmsPending)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Push sms data into mongo.
	database.DBConnection.DB("raspi_go_sms").C("sms").Insert(sms)
	// Push the sms onto the SmsQueue.
	app.SmsQueue <- sms
	fmt.Println("Sms request queued")

	// Return the sms as json.
	smsJSON, err := sms.MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(smsJSON)
	return
}

// Show function return the sms as json for the given UUID.
func (sc *SmsController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Check if the given UUID is valid.
	UUID, err := uuid.ParseHex(id)
	if err != nil {
		http.Error(w, "Malformed uuid.", http.StatusBadRequest)
		return
	}

	// Fetch the sms from mongo.
	sms := gsm.Sms{}
	if err := database.DBConnection.DB("raspi_go_sms").C("sms").FindId(UUID).One(&sms); err != nil {
		http.Error(w, "Sms not found.", http.StatusNotFound)
		return
	}

	// Return the Sms as json.
	smsJSON, err := sms.MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(smsJSON)
	return
}

// Index func
func (sc *SmsController) Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "hello")
	return
}
