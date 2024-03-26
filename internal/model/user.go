package model

import (
	"time"

	_ "gorm.io/driver/postgres"
)

type User struct {
	Id           int       `json:"-" db:"user_id"`
	Email        string    `json:"email" binding:"required" db:"email"`
	PhoneNumber  string    `json:"phone_number" binding:"reuqired" db:"phone_number"`
	Name         string    `json:"name" binding:"required" db:"name"`
	Password     string    `json:"password" binding:"required" db:"password_hash"`
	Sex          string    `json:"sex" db:"sex"`
	Age          int       `json:"age" binding:"required" db:"age"`
	BirthDate    time.Time `json:"birth_date" binding:"required" db:"birth_date"`
	City         string    `json:"city" binding:"required" db:"city"`
	Description  string    `json:"description" db:"description"`
	Role         string    `db:"role"`
	MaxAge       int       `json:"max_age" binding:"required" db:"max_age"`
	CreationDate time.Time `json:"-" db:"creation_date"`
	// Radius      int       `db:"radius"`
	// LastIP      string    `db:"last_ip"`
	// Latitude    float64   `db:"latitude"`
	// Longitude   float64   `db:"longtitude"`
}
