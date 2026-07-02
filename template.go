package main

import "html/template"

var Tmpl = template.Must(
	template.ParseFiles("templates/form.html"),
)
var Success = template.Must(
	template.ParseFiles("templates/success.html"),
)
