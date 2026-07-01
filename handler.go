package main

import (
	"html/template"
	"net/http"
)

var Tmpl = template.Must(
	template.ParseFiles("templates/index.html"),
)
var Suc = template.Must(
	template.ParseFiles("templates/sucess.html"),
)

type Ascii struct {
	Result string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		Tmpl.ExecuteTemplate(w, "index.html", nil)
		return
	}
	if r.Method == http.MethodPost {
		Text := r.FormValue("text")
		input := inputFiles(Text)
		BannerName := r.FormValue("banner")
		path := "banner/" + BannerName + ".txt"

		data, err := LoadBanner(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result := generate(input, data)
		art := Ascii{
			Result: result,
		}
		Suc.ExecuteTemplate(w, "sucess.html", art)
	}

}
