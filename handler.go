package main

import "net/http"

type Ascii struct{
	Result string
}
func homePage(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		Tmpl.ExecuteTemplate(w, "form.html", nil)
		return
	}
	if r.Method == http.MethodPost {

		text :=inputFiles(r.FormValue("text"))
		valid := map[string]bool {
			"standard": true,
			"shadow": true,
			"thinkertoy": true,
		}
		bannerName := r.FormValue("banner")
		if !valid[bannerName] {
			http.Error(w, "Invalid Banner", http.StatusBadRequest)
			return
		}
		path := "banner/" + bannerName + ".txt"
		banner, err := LoadBanner(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result := GenerateAscii(text, banner)
		art := Ascii{
			Result: result,
		}
		Success.ExecuteTemplate(w, "success.html", art)
	}
}