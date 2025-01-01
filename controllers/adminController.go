// controllers/adminController.go
package controllers

import (
    "backend/config"   // Pastikan mengimpor package config
    "backend/models"   // Pastikan mengimpor package models
    "github.com/gin-gonic/gin"
    "net/http"
    "golang.org/x/crypto/bcrypt"
)

func AdminLogin(c *gin.Context) {
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    // Bind data dari body request
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Cari admin berdasarkan email
    var user models.User
    if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Verifikasi password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Jika berhasil login
    c.JSON(http.StatusOK, gin.H{"message": "Admin login successful"})
}
