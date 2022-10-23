package controllers

import (
	"gotodo/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var CreateProject = func(c *gin.Context) {
	var project models.Project
	user, _ := c.Get("user")

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := project.Create(user.(models.User).ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

var UpdateProject = func(c *gin.Context) {
	var project models.Project
	user, _ := c.Get("user") // User Session

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := project.Update(c.Param("project"), user.(models.User).ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

var DeleteProject = func(c *gin.Context) {
	var project models.Project
	user, _ := c.Get("user") // User Session

	if err := project.Delete(c.Param("project"), user.(models.User).ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

var GetProject = func(c *gin.Context) {
	var project models.Project
	user, _ := c.Get("user") // User Session

	if err := project.Get(c.Param("project"), user.(models.User).ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

var ListProjects = func(c *gin.Context) {
	var projects models.Projects
	user, _ := c.Get("user")

	if err := projects.List(user.(models.User).ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

var ArchiveAProject = func(c *gin.Context) {
	var project models.Project
	user, _ := c.Get("user") // User Session

	if err := project.Archive(c.Param("project"), user.(models.User).ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})

}
