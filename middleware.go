package main

import (
	"net/http"
)

//AuthorizeRequest ...
func AuthorizeRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//json token may be or cookie
		//fmt.Fprintln(w, "Before autherization")
		next.ServeHTTP(w, r)
		//fmt.Fprintln(w, "After autherization")
	})
}

//JSONPayloadMiddleware ...
// func JSONPayloadMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		//set json defaults for payload
// 		fmt.Fprintln(w, "from json middlewear")
// 		next.ServeHTTP(w, r)
// 	})
// }
