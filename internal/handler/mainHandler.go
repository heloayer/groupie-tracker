package handler

import (
	"html/template"
	"net/http"
)

func (h *Handler) MainHandler(writer http.ResponseWriter, request *http.Request) {
	h.updateMain()
	if request.URL.Path != "/" {
		ErrorHandler(writer, "Page is Not Found", http.StatusNotFound)
		return
	}

	if request.Method != http.MethodGet {
		ErrorHandler(writer, "Method is not Allowed", http.StatusMethodNotAllowed)
		return
	}
	html, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err = html.Execute(writer, h); err != nil {
		ErrorHandler(writer, "Page is Not Found", http.StatusNotFound)
		return
	}
}

func (h *Handler) updateMain() {
	h.FoundBands = h.Bands
	h.Found = false
}
