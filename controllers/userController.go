package controllers

import (
	"backend/models"
	"golang.org/x/crypto/bcrypt"
	"backend/config"
	"github.com/gin-gonic/gin"
	"os"
	"fmt"
	"time"
	"path/filepath"
	"net/http"
)

func ShowSignupPage(c *gin.Context) {
	// Path menuju file signup.html
	signupPage := filepath.Join("frontends", "signup.html")

	// Cek apakah file ada
	if _, err := os.Stat(signupPage); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Signup page not found"})
		return
	}

	// Kirim file HTML sebagai respons
	c.File(signupPage)
}
func ShowAdminLoginPage(c *gin.Context) {
    // Log untuk memastikan request sampai ke controller ini
    fmt.Println("Request received at /admin/login")

    adminLoginPage := filepath.Join("frontends", "login._admin.html")
    if _, err := os.Stat(adminLoginPage); os.IsNotExist(err) {
        c.JSON(http.StatusNotFound, gin.H{"error": "Admin login page not found"})
        return
    }

    c.File(adminLoginPage)
}


// Handler untuk menampilkan halaman login
func ShowLoginPage(c *gin.Context) {
	// Path menuju file login.html
	loginPage := filepath.Join("frontends", "login.html")

	// Cek apakah file ada
	if _, err := os.Stat(loginPage); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Login page not found"})
		return
	}

	// Kirim file HTML sebagai respons
	c.File(loginPage)
}
// Fungsi SignUp untuk membuat akun pengguna baru
func SignUp(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	db := config.GetDB()
	if err := db.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user := models.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  string(hashedPassword),
		Role:      input.Role, // Pastikan ada role di input
		CreatedAt: time.Now(), // Menggunakan waktu saat ini
		UpdatedAt: time.Now(), // Menggunakan waktu saat ini
	}
	
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var input models.LoginInput
	// Bind JSON request body ke input login
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cari user berdasarkan email
	var user models.User
	db := config.GetDB()
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Periksa apakah password yang diberikan sesuai dengan yang ada di database
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Periksa apakah user adalah admin
	if user.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied, admin only"})
		return
	}

	// Jika login berhasil dan admin, buat token atau berikan akses
	// Tokenisasi atau akses diberikan di sini, bisa menggunakan JWT atau session
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}
// Fungsi EditUser untuk memperbarui data user berdasarkan ID
func EditUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	var user models.User
	db := config.GetDB()
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Name = input.Name
	user.Email = input.Email
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		user.Password = string(hashedPassword)
	}

	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// Fungsi DeleteUser untuk menghapus user berdasarkan ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	db := config.GetDB()
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// Fungsi GetUser untuk mengambil data user berdasarkan ID
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	db := config.GetDB()
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Fungsi GetAllUsers untuk mengambil seluruh data pengguna
func GetAllUsers(c *gin.Context) {
	var users []models.User
	db := config.GetDB()
	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
