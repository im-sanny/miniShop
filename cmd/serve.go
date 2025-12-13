package cmd

import (
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	initRoutes(mux)

	fmt.Println("server running on :3001")

	err := http.ListenAndServe(":3001", mux)
	// this will catch error if theres any while running the server
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
