package handler

import (
	"html/template"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles(
		"templates/about/index.html",
	))

	data := map[string]string{
		"Title":   "LinguaVerse - About",
		"Content": "Selamat datang di halaman About 🚀",
	}

	tmpl.Execute(w, data)
}