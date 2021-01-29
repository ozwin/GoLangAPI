package main

import (
	"net/http"
)

//AuthorizeRequest ...
func AuthorizeRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//json token validation can be added and request can be rejected from here if autorization fails
		next.ServeHTTP(w, r)
	})
}
