package controllers

import (
	"gotodo/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var CreateTask = func(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := task.Create(c.Param("project")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

var UpdateTask = func(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := task.Update(c.Param("task")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

var DeleteTask = func(c *gin.Context) {
	var task models.Task

	task.ProjectId = c.Param("project")
	if err := task.Delete(c.Param("task")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

var MarkAsDone = func(c *gin.Context) {
	var task models.Task

	task.ProjectId = c.Param("project")

	if err := task.MarkAsDone(c.Param("task")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}
