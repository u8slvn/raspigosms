package main

import (
	"encoding/json"
	"net/http"
)

// DataJSONResponse is the standard json response model.
type dataJSONResponse struct {
	Data  []interface{}     `json:"data"`
	Links map[string]string `json:"links,omitempty"`
	Meta  map[string]string `json:"meta,omitempty"`
}

func newDataJSONResponse(data []interface{}) dataJSONResponse {
	return dataJSONResponse{
		Data:  data,
		Links: make(map[string]string),
		Meta:  make(map[string]string),
	}
}

func (d *dataJSONResponse) addLink(key string, value string) {
	d.Links[key] = value
}

func (d *dataJSONResponse) addMeta(key string, value string) {
	d.Meta[key] = value
}

type errorJSONResponse struct {
	Errors []apiError `json:"errors"`
}

func newErrorJSONResponse() errorJSONResponse {
	return errorJSONResponse{
		Errors: []apiError{},
	}
}

func (e *errorJSONResponse) addError(apiError apiError) {
	e.Errors = append(e.Errors, apiError)
}

type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func responseJSON(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if dataJSON, err := json.Marshal(data); err == nil {
		w.Write(dataJSON)
	}
}

func responseJSONError(w http.ResponseWriter, message string, code int) {
	responseError := newErrorJSONResponse()
	responseError.addError(apiError{code, message})
	responseJSON(w, responseError, code)
}
