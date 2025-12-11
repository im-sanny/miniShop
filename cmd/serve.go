package cmd

import (
	"fmt"
	"miniShop/globalRouter"
	"miniShop/handlers"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /items", http.HandlerFunc(handlers.GetItems))
	mux.Handle("GET /items/{itemId}", http.HandlerFunc(handlers.GetItemById))
	mux.Handle("POST /items", http.HandlerFunc(handlers.PostItem))

	fmt.Println("server running on :3001")

	globalRouter := globalRouter.GlobalRouter(mux)

	err := http.ListenAndServe(":3001", globalRouter)
	// this will catch error if theres any while running the server
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
