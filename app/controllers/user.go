package controllers

import (
	"gotodo/app/models"
	"gotodo/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var UserLogin = func(c *gin.Context) {
	user := &models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.Login(user.Email, user.Password); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.CreateLoginSignature(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

var UserRegister = func(c *gin.Context) {
	user := &models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.Register(); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
