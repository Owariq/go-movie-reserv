package seatsRepository

import (
	"context"
	"fmt"
	"log"

	seatsModel "github.com/Owariq/go-movie-reserv/internal/models/seats"
	"github.com/Owariq/go-movie-reserv/internal/repository"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewSeatsRepository(db *gorm.DB) repository.SeatsRepository {
	return &repo{db: db}
}

func (r *repo) ReserveSeat(ctx context.Context, seat seatsModel.Seats) (error)  {
	var existingSeat seatsModel.Seats

	r.db.Where("seat_number = ? AND movie_id = ? AND time = ?", seat.SeatNumber, seat.MovieID, seat.Time).First(&existingSeat)
	if existingSeat.Reserved {
		return fmt.Errorf("seat already reserved")
	}

	seat.Reserved = true

	err := r.db.Create(&seat)
	if err.Error != nil {
		return fmt.Errorf("could not reserve seat")
	}

	return nil
}

func (r *repo) UnreserveSeat(ctx context.Context, seat seatsModel.Seats) error  {
	log.Printf("%+v", seat)
	r.db.Where("seat_number = ? AND movie_id = ?", seat.SeatNumber, seat.MovieID).First(&seat)
	log.Printf("%+v", seat)
	if !seat.Reserved {
		return fmt.Errorf("seat not reserved")
	}
	seat.Reserved = false
	err := r.db.Save(&seat)
	if err.Error != nil {
		return fmt.Errorf("could not unreserve seat")
	}
	return nil
}

func (r *repo) ShowReservedSeats(ctx context.Context, movieID uint) ([]seatsModel.Seats, error) {
	var reservedSeats []seatsModel.Seats
	err := r.db.Where("movie_id = ? AND reserved = ?", movieID, true).Find(&reservedSeats)
	if err.Error != nil {
		return nil, fmt.Errorf("could not show reserved seats")
	}
	return reservedSeats, nil
}