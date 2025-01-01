package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings" 
	"backend/models"
	"backend/config"
)

// Fungsi AuthMiddleware untuk memverifikasi token atau autentikasi pengguna

// AuthAdminMiddleware untuk memverifikasi apakah pengguna adalah admin
func AuthAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token atau informasi login dari header atau cookie
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// Biasanya token dikirim dalam format "Bearer <token>"
		token := strings.TrimPrefix(authHeader, "Bearer ")
		// Di sini, Anda bisa memverifikasi token dan mendapatkan data pengguna
		var user models.User
		db := config.GetDB()
		if err := db.Where("token = ?", token).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Cek apakah role adalah admin
		if user.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden, only admin can access this page"})
			c.Abort()
			return
		}

		// Simpan data user ke context untuk digunakan di handler selanjutnya
		c.Set("user", user)
		c.Next()
	}
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logika autentikasi atau pemeriksaan token (misalnya, di header)
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Misalnya validasi token (simple validation di sini)
		if token != "Bearer someValidToken" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Lanjutkan ke route berikutnya jika valid
		c.Next()
	}
}

// Fungsi untuk memverifikasi token (misalnya JWT)
func verifyToken(token string) (uint, error) {
    // Logika untuk verifikasi token dan mendapatkan user ID
    return 1, nil // Contoh: token yang valid mengembalikan ID user
}
