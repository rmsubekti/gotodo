package controllers

import (
	"gotodo/app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var CreateProject = func(c *gin.Context) {
	var project models.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := project.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

var UpdateProject = func(c *gin.Context) {
	var project models.Project
	var id int
	var err error

	if id, err = strconv.Atoi(c.Param("project")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	project.ID = uint(id)
	if err := project.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

var DeleteProject = func(c *gin.Context) {
	var project models.Project
	var id int
	var err error

	if id, err = strconv.Atoi(c.Param("project")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	project.ID = uint(id)
	if err := project.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

var GetProject = func(c *gin.Context) {
	var project models.Project
	var id int
	var err error

	if id, err = strconv.Atoi(c.Param("project")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	project.ID = uint(id)
	if err := project.Get(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

var ListProjects = func(c gin.Context) {
	var projects models.Projects

	if err := projects.List(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": projects})
}
