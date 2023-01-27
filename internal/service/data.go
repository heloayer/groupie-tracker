package service

type URLs struct {
	ArtistsUrl   string `json:"artists"`
	LocationsUrl string `json:"locations"`
	DatesUrl     string `json:"dates"`
	RelationUrl  string `json:"relation"`
}

type Searched struct {
	Artists []Band
	Found   bool
}

type Band struct {
	Id               int      `json:"id"`
	Images           string   `json:"image"`
	Name             string   `json:"name"`
	Members          []string `json:"members"`
	CreationDate     int      `json:"creationDate"`
	FirstAlbum       string   `json:"firstAlbum"`
	Locations        string   `json:"locations"`
	Dates            string   `json:"concertDates"`
	ConcertDates     []string
	ConcertLocations []string
	Relations        string `json:"relations"`
	Concerts         map[string][]string
}

type Relation struct {
	Index []struct {
		Id            int                 `json:"id"`
		DateLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}
