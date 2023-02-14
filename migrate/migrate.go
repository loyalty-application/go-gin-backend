package main

import (
	"github.com/ride-kaki/fake-rideshare-api/initializers"
	"github.com/ride-kaki/fake-rideshare-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
