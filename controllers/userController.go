package controllers

import (
	"gin/models"

	"github.com/gin-gonic/gin"
)

var users []models.User
var idCounter = 1

func GetUsers(c *gin.Context) {
	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user.Id = idCounter
	idCounter++

	users = append(users, user)
	c.JSON(201, gin.H{
		"message": "user created",
		"data":    user,
	})
}
