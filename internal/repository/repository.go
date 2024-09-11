package repository

import (
	"context"

	models "github.com/Owariq/go-movie-reserv/internal/models/auth"
	movieModels "github.com/Owariq/go-movie-reserv/internal/models/movies"
	seatsModel "github.com/Owariq/go-movie-reserv/internal/models/seats"
)

type AuthRepository interface {
	GetUser(ctx context.Context, email string) (models.User, error)
	CreateUser(ctx context.Context, user models.User) error
}

type MovieRepository interface {
	AddMovie(ctx context.Context, movie movieModels.Movie) (int, error)
	UpdateMovie(ctx context.Context, movie movieModels.Movie) error
	DeleteMovie(ctx context.Context, id uint) error
	SetSchedules(ctx context.Context, schedules movieModels.Schedules) error
	ShowSchedules(ctx context.Context, movieID uint) (movieModels.Schedules, error)
	DeleteSchedules(ctx context.Context, movieID uint) error
}

type SeatsRepository interface {
	ReserveSeat(ctx context.Context, seat seatsModel.Seats) error
	UnreserveSeat(ctx context.Context, seat seatsModel.Seats) error
	ShowReservedSeats(ctx context.Context, movieID uint) ([]seatsModel.Seats, error)
}



