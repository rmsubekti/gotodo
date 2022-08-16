package controllers

import (
	"gotodo/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var CreateProject = func(c *gin.Context) {
	var project models.Project
	var user any
	var ok bool

	if user, ok = c.Get("user"); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusInternalServerError})
		return
	}

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project.UserId = user.(models.User).ID

	if err := project.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

var UpdateProject = func(c *gin.Context) {
	var project models.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := project.Update(c.Param("project")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

var DeleteProject = func(c *gin.Context) {
	var project models.Project

	if err := project.Delete(c.Param("project")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

var GetProject = func(c *gin.Context) {
	var project models.Project

	if err := project.Get(c.Param("project")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

var ListProjects = func(c *gin.Context) {
	var projects models.Projects

	if err := projects.List(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

var ArchiveAProject = func(c *gin.Context) {
	var project models.Project

	if err := project.Archive(c.Param("project")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})

}
