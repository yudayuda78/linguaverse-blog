package main

import (
	"log"
	"net/http"

	"github.com/yudayuda78/linguaverse/internal/database"
	"github.com/yudayuda78/linguaverse/internal/handler"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/about", handler.AboutHandler)

	fileServer := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	database.Connect()

	log.Println("Server running on :3000")
	http.ListenAndServe(":3000", mux)
}
