package rest

import (
	"fmt"
	"miniShop/config"
	middleware "miniShop/rest/middlewares"
	"net/http"
	"strconv"
)

func Start(cnf config.Config) {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Cors,
		middleware.Logger,
		middleware.Preflight,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)
	initRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("server running on", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	// this will catch error if theres any while running the server
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
