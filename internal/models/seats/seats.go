package seatsModel

import (
	"time"

	"gorm.io/gorm"
)

type Seats struct {
	gorm.Model
	SeatNumber string `json:"seat_number"`
	MovieID    uint   `json:"movie_id"`
	Reserved   bool   `json:"reserved" gorm:"default:false"`
	Time 	   time.Time `json:"time"`
}