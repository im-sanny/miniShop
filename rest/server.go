package rest

import (
	"fmt"
	"miniShop/config"
	"miniShop/rest/handlers/item"
	"miniShop/rest/handlers/user"
	middleware "miniShop/rest/middlewares"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	cnf         *config.Config
	itemHandler *item.Handler
	userHandler *user.Handler
}

func NewServer(
	cnf *config.Config,
	itemHandler *item.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		cnf:         cnf,
		itemHandler: itemHandler,
		userHandler: userHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Cors,
		middleware.Logger,
		middleware.Preflight,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	server.itemHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HTTPPort)
	fmt.Println("server running on", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	// this will catch error if theres any while running the server
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
