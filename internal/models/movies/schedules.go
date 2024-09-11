package movieModels

import (
	"gorm.io/gorm"
)


type Schedules struct {
	gorm.Model
	Time    string `json:"time"`
	MovieID uint      `json:"movie_id"`
}