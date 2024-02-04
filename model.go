package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
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
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type Date struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type Relation struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

func APIHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	artists := fetchAPI("https://groupietrackers.herokuapp.com/api/artists", &[]Artist{})
	relations := fetchAPI("https://groupietrackers.herokuapp.com/api/relation", &[]Relation{})
	locations := fetchAPI("https://groupietrackers.herokuapp.com/api/locations", &[]Location{})
	dates := fetchAPI("https://groupietrackers.herokuapp.com/api/dates", &[]Date{})

	// Combine all data into a single map
	response := map[string]interface{}{
		"artists":   artists,
		"relations": relations,
		"locations": locations,
		"dates":     dates,
	}
	// Render the result in the HTML template
	templ, err := template.ParseFiles("templates/index.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Pass the result to the template for rendering
	err = templ.Execute(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func fetchAPI(apiURL string, target interface{}) interface{} {
	resp, err := http.Get(apiURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}

	return target
}
