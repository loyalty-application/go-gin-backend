package main

import (
	"github.com/cs301-itsa/project-2022-23t2-project-2022-23t2-g1-t6/initializers"
	"github.com/cs301-itsa/project-2022-23t2-project-2022-23t2-g1-t6/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
