package app

import (
	"log"
	"os"

	"github.com/Owariq/go-movie-reserv/config"
	"github.com/Owariq/go-movie-reserv/internal/api"
	"github.com/Owariq/go-movie-reserv/internal/api/movieApi"
	movieModels "github.com/Owariq/go-movie-reserv/internal/models/movies"
	"github.com/Owariq/go-movie-reserv/internal/repository"
	"github.com/Owariq/go-movie-reserv/internal/repository/movieRepository"
	"github.com/Owariq/go-movie-reserv/internal/service"
	"github.com/Owariq/go-movie-reserv/internal/service/movieService"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type movieProvider struct {
	dbConfig     config.DBConfig
	dbClient     *gorm.DB
	movieRepo    repository.MovieRepository
	movieService service.MovieService
	movieImpl    api.MovieApi
}

func NewMovieProvider() *movieProvider {
	return &movieProvider{}
}

func (p *movieProvider) DBConfig() config.DBConfig {
	p.dbConfig = config.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}
	return p.dbConfig
}

func (p *movieProvider) DBClient() *gorm.DB {
	dsn := os.Getenv("PG_DSN")

	if p.dbClient == nil {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database", err)
		}
		if err := db.AutoMigrate(&movieModels.Movie{}); err != nil {
			log.Fatal("Failed to migrate database", err)
		}
			if err := db.AutoMigrate(&movieModels.Schedules{}); err != nil {
			log.Fatal("Failed to migrate database", err)
		}
		p.dbClient = db
	}
	return p.dbClient
}

func (p *movieProvider) MovieRepo() repository.MovieRepository {
	if p.movieRepo == nil {
		p.movieRepo = movieRepository.NewMovieRepository(p.DBClient())
	}
	return p.movieRepo
}

func (p *movieProvider) MovieService() service.MovieService {
	if p.movieService == nil {
		p.movieService = movieService.NewMovieService(p.MovieRepo())
	}
	return p.movieService
}

func (p *movieProvider) MovieImpl() api.MovieApi {
	if p.movieImpl == nil {
		p.movieImpl = movieApi.NewMovieApi(p.MovieService())
	}
	return p.movieImpl
}