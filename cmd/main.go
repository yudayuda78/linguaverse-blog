package main

import (
	"log"
	"net/http"

	"github.com/yudayuda78/linguaverse/internal/handler"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/home", handler.HomeHandler)

	log.Println("Server running on :3000")
	http.ListenAndServe(":3000", mux)
}
