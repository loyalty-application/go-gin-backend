package main

import (
	"github.com/cs301-itsa/project-2022-23t2-project-2022-23t2-g1-t6/controllers"
	"github.com/cs301-itsa/project-2022-23t2-project-2022-23t2-g1-t6/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	r := gin.Default()

	// show all posts
	r.GET("/posts", controllers.PostsIndex)

	// show specific post with specific id
	r.GET("/posts/:id", controllers.PostsShow)

	// create new post
	r.POST("/posts", controllers.PostsCreate)

	// update post
	r.PUT("/posts/:id", controllers.PostsUpdate)

	// delete request
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
