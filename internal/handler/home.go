package handler

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles(
		"templates/home/index.html",
	))

	data := map[string]string{
		"Title":   "LinguaVerse - Home",
		"Content": "Selamat datang di halaman Home 🚀",
	}

	tmpl.Execute(w, data)
}
