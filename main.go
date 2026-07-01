package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(":8080", nil)

}
