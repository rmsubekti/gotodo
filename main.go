package main

import (
	"gotodo/app/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	app := route.Group("/app")
	{
		app.GET("/projects", controllers.ListProjects)
		app.GET("/projects/:project", controllers.GetProject)
		app.POST("/projects", controllers.CreateProject)
		app.PUT("/projects/:project", controllers.UpdateProject)
		app.PUT("/projects/:project/archive", controllers.ArchiveAProject)
		app.DELETE("/projects/:project", controllers.DeleteProject)

		app.POST("/projects/:project/tasks", controllers.CreateTask)
		app.POST("/projects/:project/tasks/:task", controllers.UpdateTask)
		app.PUT("/projects/:project/tasks/:task", controllers.MarkAsDone)
		app.DELETE("/projects/:project/tasks/:task", controllers.DeleteTask)
	}

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Status OK",
		})
	})
	route.Run()
}
