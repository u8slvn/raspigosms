package web

import (
	"encoding/json"
	"net/http"
)

type DataJSONResponse struct {
	Data  interface{}       `json:"data"`
	Links map[string]string `json:"links,omitempty"`
}

func NewDataJSONResponse(data interface{}) DataJSONResponse {
	return DataJSONResponse{
		Data:  data,
		Links: make(map[string]string),
	}
}

func (d *DataJSONResponse) AddLink(key string, value string) {
	d.Links[key] = value
}

type ErrorJSONResponse struct {
	Errors []APIError `json:"errors"`
}

func NewErrorJSONResponse() ErrorJSONResponse {
	return ErrorJSONResponse{
		Errors: []APIError{},
	}
}

func (e *ErrorJSONResponse) AddError(apiError APIError) {
	e.Errors = append(e.Errors, apiError)
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ResponseJSON(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if dataJSON, err := json.Marshal(data); err == nil {
		w.Write(dataJSON)
	}
}

func ResponseJSONError(w http.ResponseWriter, message string, code int) {
	responseError := NewErrorJSONResponse()
	responseError.AddError(APIError{code, message})
	ResponseJSON(w, responseError, code)
}
