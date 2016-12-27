package web

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
	phone := r.FormValue("phone")
	message := r.FormValue("message")

	sms, err := gsm.NewSms(phone, message, gsm.SmsStatusPending)
	if err != nil {
		ResponseJSONError(w, err.Error(), http.StatusNotFound)
		return
	}

	database.DBConnection.C("sms").Insert(sms)

	app.SmsRequestQueue <- app.NewSmsRequest(sms)
	fmt.Println("Sms request queued")

	responseJSON := NewDataJSONResponse([]interface{}{&sms})
	if selfPath, err := router.Get("sms_show").URLPath("id", sms.UUID.String()); err == nil {
		responseJSON.AddLink("self", selfPath.Path)
	}
	ResponseJSON(w, responseJSON, 200)
	return
}

// Show function return the sms as json for the given UUID.
func (sc *SmsController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	UUID, err := uuid.ParseHex(id)
	if err != nil {
		ResponseJSONError(w, "Malformed uuid.", http.StatusBadRequest)
		return
	}

	sms := gsm.Sms{}
	if err := database.DBConnection.C("sms").FindId(UUID).One(&sms); err != nil {
		ResponseJSONError(w, "Sms not found.", http.StatusNotFound)
		return
	}

	responseJSON := NewDataJSONResponse([]interface{}{&sms})
	if selfPath, err := router.Get("sms_show").URLPath("id", sms.UUID.String()); err == nil {
		responseJSON.AddLink("self", selfPath.Path)
	}
	ResponseJSON(w, responseJSON, 200)
	return
}

// Index func
func (sc *SmsController) Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "hello")
	return
}
