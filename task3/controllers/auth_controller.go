package controllers

import (
	"myapp/config"
	"myapp/models"
	"myapp/utils"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	user.Password = utils.HashPassword(user.Password)
	user.Role = "user"

	config.DB.Create(&user)

	c.JSON(201, gin.H{"message": "User created"})
}

func Login(c *gin.Context) {
	var input models.User
	var user models.User

	c.ShouldBindJSON(&input)

	config.DB.Where("email = ?", input.Email).First(&user)

	if !utils.CheckPassword(input.Password, user.Password) {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateToken(user.ID, user.Role)

	c.JSON(200, gin.H{"token": token})
}
