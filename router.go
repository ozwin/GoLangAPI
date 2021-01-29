package main

import (
	"github.com/gorilla/mux"
)

//CustomRouters ...
func CustomRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
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
