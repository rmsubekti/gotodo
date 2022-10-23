package main

import (
	"gotodo/app/controllers"
	middlewares "gotodo/app/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	app := route.Group("/app")
	app.Use(middlewares.Auth())
	{
		app.GET("/", controllers.ListProjects)
		app.POST("/", controllers.CreateProject)
		app.GET("/:project", controllers.GetProject)
		app.POST("/:project", controllers.UpdateProject)
		app.PUT("/:project/archive", controllers.ArchiveAProject)
		app.DELETE("/:project", controllers.DeleteProject)

		app.POST("/:project/tasks", controllers.CreateTask)
		app.POST("/:project/tasks/:task", controllers.UpdateTask)
		app.PUT("/:project/tasks/:task", controllers.MarkAsDone)
		app.DELETE("/:project/tasks/:task", controllers.DeleteTask)
	}
	auth := route.Group("/auth")
	{
		auth.POST("/login", controllers.UserLogin)
		auth.POST("/register", controllers.UserRegister)
	}
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Status OK",
		})
	})
	route.Run()
}
