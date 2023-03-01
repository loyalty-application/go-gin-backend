package main

import (
	"github.com/loyalty-application/go-gin-backend/config"
)

func main() {

	config.InitEnvironment()
	InitRoutes()

}
