package handler

import (
	"groupie-tracker/internal/service"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) SearchHandler(writer http.ResponseWriter, request *http.Request) {
	h.updateCache()
	if request.Method != http.MethodGet {
		ErrorHandler(writer, "Method is not Allowed", http.StatusMethodNotAllowed)
		return
	}
	html, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	req := request.FormValue("find")
	query := strings.Split(req, "-")

	if len(query) >= 2 {
		if !h.searchByCategory(strings.Join(query[:len(query)-1], "-"), query[len(query)-1]) {
			if len(query) == 2 {
				h.searchByCategory(service.ConvertLocation(req), "concertLocation")
			} else if len(query) == 3 {
				h.searchByCategory(service.ConvertDate(req), "concertDate")
				h.searchByCategory(service.ConvertDate(req), "firstAlbumDate")
			}
		}
	} else {
		req = strings.ToLower(req)
		h.searchAll(req)
	}
	if len(h.FoundBands) == 0 {
		h.Found = true
	}
	err = html.Execute(writer, h)
	if err != nil {
		ErrorHandler(writer, "Page is Not Found", http.StatusNotFound)
		return
	}
}

func (h *Handler) updateCache() {
	h.FoundBands = []service.Band{}
	h.Found = false
}

func (h *Handler) searchAll(query string) {
	contains := false
	for _, band := range h.Bands {
		if strings.Contains(strings.ToLower(band.Name), query) {
			h.FoundBands = append(h.FoundBands, band)
			continue
		}

		for _, artist := range band.Members {
			if strings.Contains(strings.ToLower(artist), query) {
				contains = true
				break
			}
		}
		if contains {
			h.FoundBands = append(h.FoundBands, band)
			contains = false
			continue
		}
		if strings.Contains(strconv.Itoa(band.CreationDate), query) {
			h.FoundBands = append(h.FoundBands, band)
			continue
		}
		if strings.Contains(band.FirstAlbum, query) {
			h.FoundBands = append(h.FoundBands, band)
			continue
		}
		for location, dates := range band.Concerts {
			if strings.Contains(strings.ToLower(location), query) {
				h.FoundBands = append(h.FoundBands, band)
				contains = true
				break
			}
			for _, date := range dates {
				if strings.Contains(strings.ToLower(date), query) {
					contains = true
					break
				}
			}
			if contains {
				h.FoundBands = append(h.FoundBands, band)
				contains = false
				break
			}
		}
		if contains {
			contains = false
		}
	}
}

func (h *Handler) searchByCategory(query, category string) bool {
	switch category {
	case "band":
		for _, band := range h.Bands {
			if band.Name == query {
				h.FoundBands = append(h.FoundBands, band)
			}
		}
	case "member":
		for _, band := range h.Bands {
			for _, member := range band.Members {
				if member == query {
					h.FoundBands = append(h.FoundBands, band)
				}
			}
		}
	case "creationDate":
		year, _ := strconv.Atoi(query)
		for _, band := range h.Bands {
			if band.CreationDate == year {
				h.FoundBands = append(h.FoundBands, band)
			}
		}
	case "firstAlbumDate":
		for _, band := range h.Bands {
			if band.FirstAlbum == query {
				h.FoundBands = append(h.FoundBands, band)
			}
		}
	case "concertLocation":
		for _, band := range h.Bands {
			for location := range band.Concerts {
				if location == query {
					h.FoundBands = append(h.FoundBands, band)
				}
			}
		}
	case "concertDate":
		for _, band := range h.Bands {
			for _, dates := range band.Concerts {
				for _, date := range dates {
					if date == query {
						h.FoundBands = append(h.FoundBands, band)
					}
				}
			}
		}
	default:
		return false
	}
	return true
}
