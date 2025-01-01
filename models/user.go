package models

import (
	"gorm.io/gorm"
    "time"
)

type User struct {
	gorm.Model
    ID        uint      `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    Role      string    `json:"role"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at"`
}



type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
