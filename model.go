package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"
	"fmt"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Date struct {
	ID int `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}


func fetchAPI(apiURL string, target interface{}) (interface{}, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &target)
	if err != nil {
		return nil, err
	}
	return target, nil
}

func APIHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	artists, err := fetchAPI("https://groupietrackers.herokuapp.com/api/artists", &[]Artist{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	relations, err := fetchAPI("https://groupietrackers.herokuapp.com/api/relation", &Relation{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	locations, err := fetchAPI("https://groupietrackers.herokuapp.com/api/locations", &Location{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dates, err := fetchAPI("https://groupietrackers.herokuapp.com/api/dates", &Date{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"artists":   artists,
		"relations": relations,
		"locations": locations,
		"dates":     dates,
	}

	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = templ.Execute(w, response)
if err != nil {
    fmt.Println("Error executing template:", err)
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
}

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
    parts := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(parts[len(parts)-1])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    artistInterface, err := fetchAPI("https://groupietrackers.herokuapp.com/api/artists/"+strconv.Itoa(id), &Artist{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    artist, ok := artistInterface.(*Artist)
    if !ok {
        http.Error(w, "Error type asserting artist", http.StatusInternalServerError)
        return
    }

    relationsInterface, err := fetchAPI("https://groupietrackers.herokuapp.com/api/relation/"+strconv.Itoa(id), &Relation{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    relations, ok := relationsInterface.(*Relation)
    if !ok {
        http.Error(w, "Error type asserting relations", http.StatusInternalServerError)
        return
    }

    locationsInterface, err := fetchAPI("https://groupietrackers.herokuapp.com/api/locations/"+strconv.Itoa(id), &Location{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    locations, ok := locationsInterface.(*Location)
    if !ok {
        http.Error(w, "Error type asserting locations", http.StatusInternalServerError)
        return
    }

    datesInterface, err := fetchAPI("https://groupietrackers.herokuapp.com/api/dates/"+strconv.Itoa(id), &Date{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    dates, ok := datesInterface.(*Date)
    if !ok {
        http.Error(w, "Error type asserting dates", http.StatusInternalServerError)
        return
    }

    response := map[string]interface{}{
        "artist":    artist,
        "relations": relations,
        "locations": locations,
        "dates":     dates,
    }

    templ, err := template.ParseFiles("templates/details.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    err = templ.Execute(w, response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}