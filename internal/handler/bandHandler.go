package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) BandHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		ErrorHandler(writer, "Method is not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id, ok := h.getID(request.URL.Path)
	if !ok {
		ErrorHandler(writer, "Page is Not Found", http.StatusNotFound)
		return
	}

	band := h.getBand(id)

	html, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		ErrorHandler(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = html.Execute(writer, band)
	if err != nil {
		ErrorHandler(writer, "Page is Not Found", http.StatusNotFound)
		return
	}
}

func (h *Handler) getID(urlPath string) (int, bool) {
	path := strings.Split(urlPath, "/")
	if path[len(path)-2] != "artist" {
		return 0, false
	}
	id, err := strconv.Atoi(path[len(path)-1])
	if err != nil || id < 1 || id > len(h.Bands) {
		return 0, false
	}
	return id, true
}
