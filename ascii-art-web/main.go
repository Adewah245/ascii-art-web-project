package main

import (
	"html/template"
	"net/http"
)

var Tmpl = template.Must(
	template.ParseFiles("templates/index.html"),
)

type Ascii struct {
	Text   string
	Banner string
	Result string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		Tmpl.ExecuteTemplate(w, "index.html", nil)
		return
	}
	if r.Method == http.MethodPost {
		data := Ascii{
			Text:   r.FormValue("text"),
			Banner: r.FormValue("banner"),
			Result: r.FormValue("Ascii art will appear here"),
		}
		Tmpl.ExecuteTemplate(w, "index.html", data)

	}

}

func main() {
	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(":8080", nil)

}
