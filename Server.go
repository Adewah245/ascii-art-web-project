package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)
	fmt.Println("Server: Running Successfully!.....")
	http.ListenAndServe(":8080", nil)
}