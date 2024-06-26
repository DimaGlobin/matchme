package model

import (
	"time"

	"github.com/DimaGlobin/matchme/internal/mm_errors"
	_ "gorm.io/driver/postgres"
)

const (
	male   = "male"
	female = "female"
)

type User struct {
	Id           uint64    `json:"-" db:"user_id"`
	Email        string    `json:"email" binding:"required" db:"email"`
	PhoneNumber  string    `json:"phone_number" binding:"required" db:"phone_number"`
	Name         string    `json:"name" binding:"required" db:"name"`
	Password     string    `json:"password" binding:"required" db:"password_hash"`
	Sex          string    `json:"sex" db:"sex"`
	Age          int       `json:"age" binding:"required" db:"age"`
	BirthDate    time.Time `json:"birth_date" binding:"required" db:"birth_date"`
	City         string    `json:"city" binding:"required" db:"city"`
	Description  string    `json:"description" db:"description"`
	Role         string    `json:"-" db:"role"`
	MaxAge       int       `json:"max_age" binding:"required" db:"max_age"`
	CreationDate time.Time `json:"-" db:"creation_date"`
	// Radius      int       `db:"radius"`
	// LastIP      string    `db:"last_ip"`
	// Latitude    float64   `db:"latitude"`
	// Longitude   float64   `db:"longtitude"`
}

type UserRecommendation struct {
	Id          uint64 `json:"id" db:"user_id"`
	Name        string `json:"name" binding:"required" db:"name"`
	Age         int    `json:"age" binding:"required" db:"age"`
	Description string `json:"description" db:"description"`
}

type UserInfo struct {
	Email       string `json:"email" binding:"required" db:"email"`
	PhoneNumber string `json:"phone_number" binding:"reuqired" db:"phone_number"`
	Name        string `json:"name" binding:"required" db:"name"`
	Sex         string `json:"sex" db:"sex"`
	Age         int    `json:"age" binding:"required" db:"age"`
	City        string `json:"city" binding:"required" db:"city"`
	Description string `json:"description" db:"description"`
	MaxAge      int    `json:"max_age" binding:"required" db:"max_age"`
}

type SignInBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Valid() error {
	if u.Age < 18 {
		return mm_errors.NewMmError("Invalid age")
	}

	if u.Email == "" {
		return mm_errors.NewMmError("Empty email")
	}

	if u.PhoneNumber == "" {
		return mm_errors.NewMmError("Empty phone number")
	}

	if u.BirthDate.After(time.Now().AddDate(-18, 0, 0)) {
		return mm_errors.NewMmError("Invalid birth date")
	}

	if u.MaxAge < 18 {
		return mm_errors.NewMmError("Invalid maximum age")
	}

	if u.Sex != male && u.Sex != female {
		return mm_errors.NewMmError("invalid sex")
	}

	if u.Name == "" {
		return mm_errors.NewMmError("Empty name")
	}

	return nil
}

func (s *SignInBody) Valid() error {
	if s.Email == "" {
		return mm_errors.NewMmError("Empty email")
	}

	if s.Password == "" {
		return mm_errors.NewMmError("Empty password")
	}

	return nil
}
