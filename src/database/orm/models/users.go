package models

import (
	"time"
)

type User struct {
	UserId       string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"user_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Gender       string `json:"gender"`
	Address      string `json:"address"`
	MobileNumber string `json:"mobile_number"`
	DisplayName  string `json:"display_name"`
	BirthDate    string `json:"birth_date"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Users []User
