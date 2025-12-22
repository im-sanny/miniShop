package cmd

import (
	"fmt"
	"miniShop/config"
	"miniShop/infra/db"
	"miniShop/item"
	"miniShop/repo"
	"miniShop/rest"
	itemHandler "miniShop/rest/handlers/item"
	userHandler "miniShop/rest/handlers/user"
	middleware "miniShop/rest/middlewares"
	"miniShop/user"
)

// - NewServer builds the server with everything it needs,
// - Server.Start() runs it,
// - Serve() is just the orchestrator.

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
	}

	// repos
	itemRepo := repo.NewItemRepo(*dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	// domains
	userSvc := user.NewService(userRepo)
	itemSvc := item.NewService(itemRepo)

	// handlers
	middlewares := middleware.NewMiddlewares(cnf)
	itemHandler := itemHandler.NewHandler(middlewares, itemSvc)
	userHandler := userHandler.NewHandler(cnf, userSvc)

	server := rest.NewServer(cnf, itemHandler, userHandler)
	server.Start()
}
