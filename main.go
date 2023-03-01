package main

import (
	"github.com/loyalty-application/go-gin-backend/config"
	"github.com/loyalty-application/go-gin-backend/routes"
)

func main() {

	config.InitEnvironment()
	config.InitDatabase()
	routes.InitRoutes()

}
