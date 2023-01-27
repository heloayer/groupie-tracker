package service

import "testing"

func TestGetData(t *testing.T) {
	bands := []Band{}
	url := "https://groupietrackers.herokuapp.com/api/artists"
	GetData(url, &bands)
	if bands[0].Name != "Queen" {
		t.Errorf("got %s, wanted 'Queen'", bands[0].Name)
	}
	if bands[2].CreationDate != 1965 {
		t.Errorf("got %d, wanted '1965'", bands[2].CreationDate)
	}
	if bands[4].FirstAlbum != "07-09-2017" {
		t.Errorf("got %s, wanted '07-09-2017'", bands[4].FirstAlbum)
	}
	if bands[6].Members[0] != "Gary Maurice Lucas Jr." {
		t.Errorf("got %s, wanted 'Gary Maurice Lucas Jr.'", bands[6].Members[0])
	}
	if bands[9].Name != "Pearl Jam" {
		t.Errorf("got %s, wanted 'Pearl Jam'", bands[9].Name)
	}
}
