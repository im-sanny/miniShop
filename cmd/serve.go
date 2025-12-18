package cmd

import (
	"miniShop/config"
	"miniShop/rest"
	"miniShop/rest/handlers/item"
	"miniShop/rest/handlers/user"
	middleware "miniShop/rest/middlewares"
)

// - NewServer builds the server with everything it needs,
// - Server.Start() runs it,
// - Serve() is just the orchestrator.

func Serve() {
	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cnf)
	itemHandler := item.NewHandler(middlewares)
	userHandler := user.NewHandler()

	server := rest.NewServer(cnf, itemHandler, userHandler)
	server.Start()
}
