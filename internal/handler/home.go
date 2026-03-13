package handler

import (
	"fmt"
	"github.com/yudayuda78/linguaverse/internal/repository"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler: HomeHandler reached")
	posts, err := repository.GetAllPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"templates/home/index.html",
	))

	data := map[string]interface{}{
		"Title":   "LinguaVerse - Home",
		"Content": "Selamat datang di halaman Home 🚀",
		"Posts":   posts,
	}

	fmt.Printf("Handler: Rendering home with %d posts\n", len(posts))
	tmpl.Execute(w, data)
}
