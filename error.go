package system

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type Go2Validator interface {
	Validate() url.Values
}

func ThrowValidationError(w http.ResponseWriter, t Go2Validator) {
	if validErrs := t.Validate(); len(validErrs) > 0 {
		err := map[string]interface{}{"error": validErrs}
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		panic("Error ")
	}
}

func ThrowHttpError(w http.ResponseWriter, message string, statusCode int) {
	errorJSON := jsonErr{Code: statusCode, Text: message}
	err := map[string]interface{}{"error": errorJSON}
	w.Header().Set("Content-type", "applciation/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(err)
	panic("Error")
}
