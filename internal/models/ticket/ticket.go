package ticketModel

import (
	movieModels "github.com/Owariq/go-movie-reserv/internal/models/movies"
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Movie    movieModels.Movie   `json:"movie"`
	SeatNumber string `json:"seat_number"`
	Reserved   bool   `json:"reserved" gorm:"default:false"`
}


