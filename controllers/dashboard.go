package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Fungsi Dashboard untuk menampilkan data hanya untuk admin
func Dashboard(c *gin.Context) {
	// Ambil data user yang sudah diset dalam context
	user, _ := c.Get("user")

	// Menampilkan data dashboard untuk admin
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Admin Dashboard!",
		"user":    user,
	})
}
