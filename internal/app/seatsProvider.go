package app

import (
	"log"
	"os"

	"github.com/Owariq/go-movie-reserv/config"
	"github.com/Owariq/go-movie-reserv/internal/api"
	"github.com/Owariq/go-movie-reserv/internal/api/seatsApi"
	seatsModel "github.com/Owariq/go-movie-reserv/internal/models/seats"
	"github.com/Owariq/go-movie-reserv/internal/repository"
	"github.com/Owariq/go-movie-reserv/internal/repository/seatsRepository"
	"github.com/Owariq/go-movie-reserv/internal/service"
	"github.com/Owariq/go-movie-reserv/internal/service/seatsService"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type seatsProvider struct {
	dbConfig     config.DBConfig
	dbClient     *gorm.DB
	seatsRepo    repository.SeatsRepository
	seatsService service.SeatsService
	seatsImpl    api.SeatsApi
}

func NewSeatsProvider() *seatsProvider {
	return &seatsProvider{}
}

func (p *seatsProvider) DBConfig() config.DBConfig {
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

func (p *seatsProvider) DBClient() *gorm.DB {
	dsn := os.Getenv("PG_DSN")

	if p.dbClient == nil {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database", err)
		}
		if err := db.AutoMigrate(&seatsModel.Seats{}); err != nil {
			log.Fatal("Failed to migrate database", err)
		}
		p.dbClient = db
	}
	return p.dbClient
}

func (p *seatsProvider) SeatsRepo() repository.SeatsRepository {
	if p.seatsRepo == nil {
		p.seatsRepo = seatsRepository.NewSeatsRepository(p.DBClient())
	}
	return p.seatsRepo
}

func (p *seatsProvider) SeatsService() service.SeatsService {
	if p.seatsService == nil {
		p.seatsService = seatsService.NewSeatsService(p.SeatsRepo())
	}
	return p.seatsService
}

func (p *seatsProvider) SeatsImpl() api.SeatsApi {
	if p.seatsImpl == nil {
		p.seatsImpl = seatsApi.NewSeatsApi(p.SeatsService())
	}
	return p.seatsImpl
}