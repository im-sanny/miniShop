package cmd

import (
	"miniShop/config"
	"miniShop/rest"
)

func Serve() {
	cnf := config.GetConfig()
	rest.Start(cnf)
}
