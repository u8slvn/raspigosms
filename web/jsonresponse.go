package web

import (
	"encoding/json"
	"net/http"
)

// DataJSONResponse is the standard json response model.
type DataJSONResponse struct {
	Data  []interface{}     `json:"data"`
	Links map[string]string `json:"links,omitempty"`
	Meta  map[string]string `json:"meta,omitempty"`
}

// NewDataJSONResponse create and return a default json response struct.
// Only Data is mandatory, Links and Meta can be feed by theirs own setters.
func NewDataJSONResponse(data []interface{}) DataJSONResponse {
	return DataJSONResponse{
		Data:  data,
		Links: make(map[string]string),
		Meta:  make(map[string]string),
	}
}

// AddLink allows to feed the DataJSONResponse's Links.
func (d *DataJSONResponse) AddLink(key string, value string) {
	d.Links[key] = value
}

// AddMeta allows to feed the DataJSONResponse's Meta.
func (d *DataJSONResponse) AddMeta(key string, value string) {
	d.Meta[key] = value
}

// ErrorJSONResponse is the standard json error response model.
type ErrorJSONResponse struct {
	Errors []APIError `json:"errors"`
}

// NewErrorJSONResponse init and return a default json error response.
func NewErrorJSONResponse() ErrorJSONResponse {
	return ErrorJSONResponse{
		Errors: []APIError{},
	}
}

// AddError allows to feed the ErrorJSONResponse's Errors.
func (e *ErrorJSONResponse) AddError(apiError APIError) {
	e.Errors = append(e.Errors, apiError)
}

// APIError model
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ResponseJSON correctly set an http.RresponeWriter with data and http status to return.
func ResponseJSON(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if dataJSON, err := json.Marshal(data); err == nil {
		w.Write(dataJSON)
	}
}

// ResponseJSONError allows to return just one json api error.
func ResponseJSONError(w http.ResponseWriter, message string, code int) {
	responseError := NewErrorJSONResponse()
	responseError.AddError(APIError{code, message})
	ResponseJSON(w, responseError, code)
}
