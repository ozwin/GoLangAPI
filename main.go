package main

import (
	"net/http"
)

func main() {
	router := CustomRouter()
	http.ListenAndServe(":8080", router)
}
