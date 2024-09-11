package movieService

import (
	"context"

	movieModels "github.com/Owariq/go-movie-reserv/internal/models/movies"
	"github.com/Owariq/go-movie-reserv/internal/repository"
	"github.com/Owariq/go-movie-reserv/internal/service"
)

type serv struct {
	repo repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) service.MovieService {
	return &serv{repo: repo}
}

func (s *serv) AddMovie(ctx context.Context, movie movieModels.Movie) (int, error) {
	return s.repo.AddMovie(ctx, movie)
}

func (s *serv) UpdateMovie(ctx context.Context, movie movieModels.Movie) error {
	return s.repo.UpdateMovie(ctx, movie)
}

func (s *serv) DeleteMovie(ctx context.Context, id uint) error {
	return s.repo.DeleteMovie(ctx, id)
}

func (s *serv) SetSchedules(ctx context.Context, schedules movieModels.Schedules) error {
	return s.repo.SetSchedules(ctx, schedules)
}
func (s *serv) ShowSchedules(ctx context.Context, movieID uint) (movieModels.Schedules, error) {
	return s.repo.ShowSchedules(ctx, movieID)
}
func (s *serv) DeleteSchedules(ctx context.Context, movieID uint) error {
	return s.repo.DeleteSchedules(ctx, movieID)
}