package controllers

import (
	"gotodo/app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var CreateTask = func(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := task.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

var UpdateTask = func(c *gin.Context) {
	var task models.Task
	var id int
	var err error

	if id, err = strconv.Atoi(c.Param("task")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	task.ID = uint(id)
	if err := task.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

var DeleteTask = func(c *gin.Context) {
	var task models.Task
	var projectID int
	var id int
	var err error

	if projectID, err = strconv.Atoi(c.Param("task")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if id, err = strconv.Atoi(c.Param("task")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	task.ID = uint(id)
	task.ProjectID = uint(projectID)
	if err := task.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

var MarkAsDone = func(c *gin.Context) {
	var task models.Task
	var projectID int
	var id int
	var err error

	if projectID, err = strconv.Atoi(c.Param("task")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if id, err = strconv.Atoi(c.Param("task")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	task.ID = uint(id)
	task.ProjectID = uint(projectID)

	if err := task.MarkAsDone(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}
