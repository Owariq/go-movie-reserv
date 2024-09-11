package seatsService

import (
	"context"

	seatsModel "github.com/Owariq/go-movie-reserv/internal/models/seats"
	"github.com/Owariq/go-movie-reserv/internal/repository"
	"github.com/Owariq/go-movie-reserv/internal/service"
)

type serv struct {
	repo repository.SeatsRepository
}

func NewSeatsService(repo repository.SeatsRepository) service.SeatsService {
	return &serv{repo: repo}
}

	func (s *serv) ReserveSeat(ctx context.Context, seat seatsModel.Seats) error {
		return s.repo.ReserveSeat(ctx, seat)
	}

	func (s *serv) UnreserveSeat(ctx context.Context, seat seatsModel.Seats) error {
		return s.repo.UnreserveSeat(ctx, seat)
	}

	func (s *serv) ShowReservedSeats(ctx context.Context, movieID uint) ([]seatsModel.Seats, error) {
		return s.repo.ShowReservedSeats(ctx, movieID)
	}
