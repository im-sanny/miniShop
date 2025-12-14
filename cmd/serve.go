package cmd

import (
	"fmt"
	"miniShop/middleware"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Cors,
		middleware.Logger,
		middleware.Preflight,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)
	initRoutes(mux, manager)

	fmt.Println("server running on :3001")

	err := http.ListenAndServe(":3001", wrappedMux)
	// this will catch error if theres any while running the server
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
