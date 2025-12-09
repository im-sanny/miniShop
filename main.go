package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// album represents data about a record album
type album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "give me get req", 400)
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(albums)
}

func postAlbum(w http.ResponseWriter, r *http.Request) {
	var newAlbum album

	if r.Method != http.MethodPost {
		http.Error(w, "give proper req", 400)
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newAlbum)
	if err != nil {
		fmt.Println(err)
	}

	newAlbum.ID = len(albums) + 1

	albums = append(albums, newAlbum)
	encoder := json.NewEncoder(w)
	encoder.Encode(albums)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/albums", getAlbums)
	mux.HandleFunc("/create", postAlbum)

	err := http.ListenAndServe(":3001", mux)
	// this will catch error if theres any while running the server
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
