package main

import (
	"net/http"
)

func main() {
	router := CustomRouters()
	http.ListenAndServe(":8080", router)
}
