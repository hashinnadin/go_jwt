package routes

import (
	"myapp/controllers"
	"myapp/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminOnly())

	admin.GET("/users", controllers.GetUsers)
	admin.POST("/users", controllers.CreateUser)
	admin.PUT("/users/:id", controllers.UpdateUser)
	admin.DELETE("/users/:id", controllers.DeleteUser)
}
