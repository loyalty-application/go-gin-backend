package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ride-kaki/fake-rideshare-api/initializers"
	"github.com/ride-kaki/fake-rideshare-api/models"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		// get data of request body
		Body  string
		Title string
	}

	c.Bind(&body)

	// create a Post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func PostsIndex(c *gin.Context) {
	// get posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// get id off url
	id := c.Param("id")

	// get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// respond
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// get id off url
	id := c.Param("id")

	// get data of request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// find the post
	var post models.Post
	initializers.DB.First(&post, id)

	// update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// respond with it
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsDelete(c *gin.Context) {

	// get the id param
	id := c.Param("id")

	// delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// respond
	c.Status(200)

}
