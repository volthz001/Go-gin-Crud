package routes

import (
    "backend/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
    router.POST("/signup", controllers.SignUp)
    router.POST("/login", controllers.Login)
    router.PUT("/users/:id", controllers.EditUser)
    router.DELETE("/users/:id", controllers.DeleteUser)
	router.GET("/admin/login", controllers.ShowAdminLoginPage)
		
	
	// Endpoint untuk melihat data pengguna berdasarkan ID
	router.GET("/users/:id", controllers.GetUser)
	// Endpoint untuk melihat semua pengguna
	router.GET("/users", controllers.GetAllUsers)
}
