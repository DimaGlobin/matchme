package model

import (
	"time"

	"github.com/DimaGlobin/matchme/internal/mm_errors"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	male   = "male"
	female = "female"
)

type User struct {
	Id           uint64    `json:"-" gorm:"primaryKey;column:user_id"`
	Email        string    `json:"email" binding:"required" gorm:"column:email"`
	PhoneNumber  string    `json:"phone_number" binding:"required" gorm:"column:phone_number"`
	Name         string    `json:"name" binding:"required" gorm:"column:name"`
	Password     string    `json:"password" binding:"required" gorm:"column:password_hash"`
	Sex          string    `json:"sex" gorm:"column:sex"`
	Age          int       `json:"age" binding:"required" gorm:"column:age"`
	BirthDate    time.Time `json:"birth_date" binding:"required" gorm:"column:birth_date"`
	City         string    `json:"city" binding:"required" gorm:"column:city"`
	Description  string    `json:"description" gorm:"column:description"`
	Role         string    `json:"-" gorm:"column:role"`
	MaxAge       int       `json:"max_age" binding:"required" gorm:"column:max_age"`
	CreationDate time.Time `json:"-" gorm:"column:creation_date"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.Role == "" {
		u.Role = "basic"
	}
	return nil
}

type UserRecommendation struct {
	Id          uint64 `json:"id" gorm:"primaryKey;column:user_id"`
	Name        string `json:"name" binding:"required" gorm:"column:name"`
	Age         int    `json:"age" binding:"required" gorm:"column:age"`
	Description string `json:"description" gorm:"column:description"`
}

type UserInfo struct {
	Email       string `json:"email" binding:"required" gorm:"column:email"`
	PhoneNumber string `json:"phone_number" binding:"required" gorm:"column:phone_number"`
	Name        string `json:"name" binding:"required" gorm:"column:name"`
	Sex         string `json:"sex" gorm:"column:sex"`
	Age         int    `json:"age" binding:"required" gorm:"column:age"`
	City        string `json:"city" binding:"required" gorm:"column:city"`
	Description string `json:"description" gorm:"column:description"`
	MaxAge      int    `json:"max_age" binding:"required" gorm:"column:max_age"`
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
