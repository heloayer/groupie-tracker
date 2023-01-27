package handler

import (
	"html/template"
	"net/http"
)

type ErrorBody struct {
	Message string
	Code    int
}

func ErrorHandler(writer http.ResponseWriter, message string, code int) {
	writer.WriteHeader(code)
	errorBody := ErrorBody{Message: message, Code: code}
	html, err := template.ParseFiles("templates/errors.html")
	if err != nil {
		http.Error(writer, "500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err = html.Execute(writer, errorBody); err != nil {
		http.Error(writer, "404: Not Found", 404)
		return
	}
}
