package main

import (
	"net/http"
	"html/template"
)

var Tmpl = template.Must(
	template.ParseFiles("templates/index.html"),
)


type Page struct{
	Text string
	Banner string
}
func homeHandler(w http.ResponseWriter, r *http.Request){
	Tmpl.ExecuteTemplate(w, "index.html", nil)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{
		Text: r.FormValue("text"),
		Banner: r.FormValue("banner"),
	}
	Tmpl.ExecuteTemplate(w, "index.html", page)
}


func main() {
	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(":8080", nil)

}