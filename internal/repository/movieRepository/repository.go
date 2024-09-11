package movieRepository

import (
	"context"
	"fmt"
	"log"

	movieModels "github.com/Owariq/go-movie-reserv/internal/models/movies"
	"github.com/Owariq/go-movie-reserv/internal/repository"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) repository.MovieRepository {
	return &repo{db: db}
}

func (r *repo) AddMovie(ctx context.Context, movie movieModels.Movie) (int, error) {
	
	r.db.Create(&movie)
	
	return int(movie.ID), nil
}

func (r *repo) UpdateMovie(ctx context.Context, movie movieModels.Movie) error {
	var unUpdatedMovie movieModels.Movie
	r.db.First(&unUpdatedMovie, movie.ID)
	if unUpdatedMovie.ID == 0 {
		return fmt.Errorf("movie does not exist")
	}
	unUpdatedMovie = movie

	err := r.db.Save(&unUpdatedMovie)
	if err.Error != nil {
		return fmt.Errorf("could not update movie")
	}
	return nil
}

func (r *repo) DeleteMovie(ctx context.Context, id uint) error {
	err := r.db.Where("id = ?", id).Delete(&movieModels.Movie{})
	if err.Error != nil {
		return fmt.Errorf("could not delete movie")
	}

	return nil
}

	func (r *repo) SetSchedules(ctx context.Context, schedules movieModels.Schedules) error {
		log.Printf("%+v", schedules)
		err := r.db.Create(&schedules)
		if err.Error != nil {
			return fmt.Errorf("could not set schedules")
		}
		return nil
	}


func (r *repo) ShowSchedules(ctx context.Context, movieID uint) (movieModels.Schedules, error) {

	var schedules movieModels.Schedules
	err := r.db.Where("movie_id = ?", movieID).Find(&schedules)
	if err.Error != nil {
		return movieModels.Schedules{}, fmt.Errorf("could not show schedules")
	}
	return schedules, nil
}
func (r *repo) DeleteSchedules(ctx context.Context, movieID uint) error {

	err := r.db.Where("movie_id = ?", movieID).Delete(&movieModels.Schedules{})
	if err.Error != nil {
		return fmt.Errorf("could not delete schedules")
	}
	return nil
}