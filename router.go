package main

import (
	"github.com/gorilla/mux"
)

//CustomRouter ...
func CustomRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	//multiple middleware can be registered here
	router.Use(AuthorizeRequest)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
