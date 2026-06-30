package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var Tmpl = template.Must(
	template.ParseFiles("templates/index.html"),
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
	Tmpl.ExecuteTemplate(w, "index.html", nil)
		return
	}
	if r.Method == http.MethodPost{
		Text := r.FormValue("text")
		banner := r.FormValue("banner")
		fmt.Println("text:", Text)
		fmt.Println("banner:", banner)
	}
	Tmpl.ExecuteTemplate(w, "index.html", nil)
}

func main() {
	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(":8080", nil)

}