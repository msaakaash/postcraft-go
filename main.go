package main

import (
	"go-crud/controllers"
	"go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	//POST REQUEST
	r.POST("/posts", controllers.PostsCreate)

	//GET REQUEST
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)

	//UPDATE REQUEST
	r.PUT("/posts/:id", controllers.PostsUpdate)

	//DELETE REQUEST
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run()
}
