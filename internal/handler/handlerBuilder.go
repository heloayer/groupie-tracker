package handler

import (
	"groupie-tracker/internal/service"
)

type Handler struct {
	Bands      []service.Band
	relations  service.Relation
	FoundBands []service.Band
	Found      bool
}

func BuildHandler() Handler {
	var h Handler
	h.parseBands()
	return h
}

func (h *Handler) parseBands() {
	urls := service.URLs{}
	service.GetData("https://groupietrackers.herokuapp.com/api", &urls)
	service.GetData(urls.ArtistsUrl, &h.Bands)
	for ix := 0; ix < len(h.Bands); ix++ {
		h.Bands[ix].FirstAlbum = service.ConvertDate(h.Bands[ix].FirstAlbum)
	}
	service.GetData(urls.RelationUrl, &h.relations)
	h.addRelations()
}

func (h *Handler) getBand(id int) service.Band {
	return h.Bands[id-1]
}

func (h *Handler) addRelations() {
	for i := 0; i < len(h.Bands); i++ {
		h.Bands[i].Concerts = map[string][]string{}
		for loc, dates := range h.relations.Index[i].DateLocations {
			loc = service.ConvertLocation(loc)
			h.Bands[i].Concerts[loc] = []string{}
			for _, date := range dates {
				h.Bands[i].Concerts[loc] = append(h.Bands[i].Concerts[loc], service.ConvertDate(date))
			}
		}
	}
}
