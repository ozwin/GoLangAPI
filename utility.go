package main

import (
	"encoding/json"
	"net/http"
)

//JSONResponseWriter ...
func JSONResponseWriter(w http.ResponseWriter, payload interface{}, httpStatusCode int) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	w.Write(response)
}
