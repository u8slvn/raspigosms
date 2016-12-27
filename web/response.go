package web

import (
	"encoding/json"
	"net/http"
)

type DataResponse struct {
	Data  interface{}       `json:"data"`
	Links map[string]string `json:"links"`
}

type ErrorResponse struct {
	Errors APIError `json:"errors"`
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
	apiError := APIError{
		Code:    code,
		Message: message,
	}

	ResponseJSON(w, ErrorResponse{apiError}, code)
}
