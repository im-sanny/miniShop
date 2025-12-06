package main

import (
	"fmt"
	"net/http"
)

func shopHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Mini Shop!")
}

func ownerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Theo Fumis")
}

func main() {
	http.HandleFunc("/", shopHandler)
	http.HandleFunc("/owner", ownerHandler)

	err := http.ListenAndServe(":3001", nil)
	// this will catch error if theres any while running the server
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
