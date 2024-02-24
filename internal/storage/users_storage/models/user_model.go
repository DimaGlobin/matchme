package storage

import (
	"time"

	"github.com/lib/pq"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string `gorm:"unique"`
	Password    string
	Sex         string
	Age         int
	BirthDate   time.Time
	Location    string
	Description string `gorm:"type:TEXT"`
	Rights      string
	MaxAge      int
	Radius      int
	Liked       pq.Int64Array `gorm:"type:bigint[]"`
	Disliked    pq.Int64Array `gorm:"type:bigint[]"`
	Matches     pq.Int64Array `gorm:"type:bigint[]"`
	LastIP      string
	Latitude    float64
	Longitude   float64
}
