package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/u8slvn/raspigosms/raspigosms"
)

type smsHandler struct{}

func newSmsHandler() *smsHandler {
	return &smsHandler{}
}

func (sc *smsHandler) create(w http.ResponseWriter, r *http.Request) {
	phone := r.FormValue("phone")
	message := r.FormValue("message")

	fmt.Println("Sms request queued")
	sms, err := raspigosms.SubmitSms(phone, message)
	if err != nil {
		responseJSONError(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonResponse := newDataJSONResponse([]interface{}{&sms})
	if selfPath, err := router.Get("sms_show").URLPath("id", sms.UUID.String()); err == nil {
		jsonResponse.addLink("self", selfPath.Path)
	}
	responseJSON(w, jsonResponse, 200)
	return
}

func (sc *smsHandler) show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	sms, err := raspigosms.FindSmsByUUID(id)
	if err != nil {
		responseJSONError(w, "Sms not found.", http.StatusNotFound)
		return
	}

	jsonResponse := newDataJSONResponse([]interface{}{&sms})
	if selfPath, err := router.Get("sms_show").URLPath("id", sms.UUID.String()); err == nil {
		jsonResponse.addLink("self", selfPath.Path)
	}
	responseJSON(w, jsonResponse, 200)
	return
}

func (sc *smsHandler) index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "hello")
	return
}
