package controllers

import (
	"fmt"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/u8slvn/raspigosms/gsm"
)

// SmsQueue is a buffered channel used to send sms requests on.
var SmsQueue = make(chan gsm.Sms, 100)

type SmsController struct {
	session *mgo.Session
}

// NewSmsController creates and return a new SmsController.
func NewSmsController(session *mgo.Session) *SmsController {
	return &SmsController{session}
}

// Create collect post data to create Sms wich sent onto the SmsQueue channel.
func (sc *SmsController) Create(w http.ResponseWriter, r *http.Request) {
	// Retrieve the sms information from the request.
	phone := r.FormValue("phone")
	message := r.FormValue("message")

	// Try to create the sms.
	sms, err := gsm.NewSms(phone, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Push the sms onto the SmsQueue.
	SmsQueue <- sms
	fmt.Println("Sms request queued")

	smsToJSON, err := sms.MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(smsToJSON)
	return
}

// Index func
func (sc *SmsController) Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "hello")
	return
}
