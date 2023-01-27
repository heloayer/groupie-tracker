package handler

import (
	"fmt"
	"groupie-tracker/internal/handler"
	"log"
	"net/http"
)

func Run() {
	h := handler.BuildHandler()
	http.HandleFunc("/", h.MainHandler)
	http.HandleFunc("/artist/", h.BandHandler)
	http.HandleFunc("/search", h.SearchHandler)
	http.Handle("/templates/css/", http.StripPrefix("/templates/css/", http.FileServer(http.Dir("templates/css"))))
	http.Handle("/templates/img/", http.StripPrefix("/templates/img/", http.FileServer(http.Dir("templates/img"))))
	fmt.Printf("Starting server at post: 8087\nhttp://localhost:8087/\n")
	log.Fatal(http.ListenAndServe(":8087", nil))
}
