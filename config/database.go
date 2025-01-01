// config/config.go
package config

import (
    "backend/models"
    "fmt"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB // Deklarasikan DB sebagai variabel global

// Fungsi untuk menghubungkan ke database
func ConnectDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Auto-migrate schema, pastikan model User sudah didefinisikan
    err = DB.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

    fmt.Println("Database connected successfully!")
}

// Fungsi untuk mendapatkan koneksi database
func GetDB() *gorm.DB {
    return DB
}