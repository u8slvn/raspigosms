package web

import (
	"fmt"
	"net/http"

	"github.com/u8slvn/raspigosms/gsm"
)

// SmsQueue : A buffered channel that we can send work requests on.
var SmsQueue = make(chan gsm.Sms, 100)

// Collector function
func Collector(w http.ResponseWriter, r *http.Request) {
	// Retrieve the sms information from the request.
	phone := r.FormValue("phone")
	message := r.FormValue("message")

	// Try to create the sms.
	sms, err := gsm.NewSms(phone, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Push the sms onto the queue.
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
