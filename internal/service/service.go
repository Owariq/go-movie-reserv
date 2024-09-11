package service

import (
	"context"

	models "github.com/Owariq/go-movie-reserv/internal/models/auth"
	movieModels "github.com/Owariq/go-movie-reserv/internal/models/movies"
	seatsModel "github.com/Owariq/go-movie-reserv/internal/models/seats"
)



type AuthService interface {
	Login(ctx context.Context, email, password string) (string, int, error)
	Signup(ctx context.Context, user models.User) error
	// Home()
	// Logout()
	// Admin()
}

type MovieService interface {
	AddMovie(ctx context.Context, movie movieModels.Movie) (int, error)
	UpdateMovie(ctx context.Context, movie movieModels.Movie) error
	DeleteMovie(ctx context.Context, id uint) error
	SetSchedules(ctx context.Context, schedules movieModels.Schedules) error
	ShowSchedules(ctx context.Context, movieID uint) (movieModels.Schedules, error)
	DeleteSchedules(ctx context.Context, movieID uint) error
}

type SeatsService interface {
	ReserveSeat(ctx context.Context, seat seatsModel.Seats) error
	UnreserveSeat(ctx context.Context, seat seatsModel.Seats) error
	ShowReservedSeats(ctx context.Context, movieID uint) ([]seatsModel.Seats, error)
}

