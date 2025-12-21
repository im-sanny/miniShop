package cmd

import (
	"fmt"
	"miniShop/config"
	"miniShop/infra/db"
	"miniShop/repo"
	"miniShop/rest"
	"miniShop/rest/handlers/item"
	"miniShop/rest/handlers/user"
	middleware "miniShop/rest/middlewares"
	"os"
)

// - NewServer builds the server with everything it needs,
// - Server.Start() runs it,
// - Serve() is just the orchestrator.

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	itemRepo := repo.NewItemRepo(*dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	middlewares := middleware.NewMiddlewares(cnf)

	itemHandler := item.NewHandler(middlewares, itemRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(cnf, itemHandler, userHandler)
	server.Start()
}
