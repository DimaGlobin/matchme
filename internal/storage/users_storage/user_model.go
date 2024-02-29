package users_storage

import (
	"time"

	_ "gorm.io/driver/postgres"
)

type User struct {
	Id          int       `db:"id"`
	Email       string    `db:"email"`
	PhoneNumber string    `db:"phone_number"`
	Name        string    `db:"name"`
	Password    string    `db:"password_hash"`
	Sex         string    `db:"sex"`
	Age         int       `db:"age"`
	BirthDate   time.Time `db:"age"`
	City        string    `db:"city"`
	Description string    `db:"description"`
	Role        string    `db:"role"`
	MaxAge      int       `db:"max_age"`
	// Radius      int       `db:"radius"`
	// LastIP      string    `db:"last_ip"`
	// Latitude    float64   `db:"latitude"`
	// Longitude   float64   `db:"longtitude"`
}
