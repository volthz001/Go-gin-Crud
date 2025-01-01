package main

import (
    "backend/config"
    "backend/controllers"
    "backend/middleware"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    // Koneksi ke database
    config.ConnectDB()

    // Inisialisasi router
    r := gin.Default()

    // Menambahkan middleware CORS untuk mengizinkan akses dari semua origin
    r.Use(cors.New(cors.Config{
        AllowOrigins: []string{"http://127.0.0.1:8080"},
        AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
    }))

    // Rute publik
    r.GET("/signup", controllers.ShowSignupPage)
    r.POST("/signup", controllers.SignUp)

    r.GET("/login", controllers.ShowLoginPage)
    r.POST("/login", controllers.Login)

    // Rute untuk login admin
    r.POST("/admin/login", controllers.AdminLogin)

    // Rute yang membutuhkan autentikasi admin
    admin := r.Group("/admin")
    admin.Use(middleware.AuthAdminMiddleware())
    {
        admin.GET("/dashboard", controllers.Dashboard)
    }

    // Rute yang membutuhkan autentikasi pengguna
    protected := r.Group("/users")
    protected.Use(middleware.AuthMiddleware())
    {
        protected.PUT("/:id", controllers.EditUser)
        protected.DELETE("/:id", controllers.DeleteUser)
    }

    // Jalankan server
    r.Run(":8080")
}
