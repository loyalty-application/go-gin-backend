package main

import (
	"github.com/loyalty-application/go-gin-backend/config"
	"github.com/loyalty-application/go-gin-backend/db"
	"github.com/loyalty-application/go-gin-backend/server"
)

func main() {

	config.Init()
	db.Init()
	server.Init()

}
