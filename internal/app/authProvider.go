package app

import (
	"log"
	"os"

	"github.com/Owariq/go-movie-reserv/config"
	"github.com/Owariq/go-movie-reserv/internal/api"
	"github.com/Owariq/go-movie-reserv/internal/api/authApi"
	models "github.com/Owariq/go-movie-reserv/internal/models/auth"
	"github.com/Owariq/go-movie-reserv/internal/repository"
	"github.com/Owariq/go-movie-reserv/internal/repository/authRepository"
	"github.com/Owariq/go-movie-reserv/internal/service"
	"github.com/Owariq/go-movie-reserv/internal/service/authService"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type authProvider struct {
	dbConfig config.DBConfig
	dbClient *gorm.DB
	authRepo repository.AuthRepository
	authService service.AuthService
	authImpl api.AuthApi
}

func NewAuthProvider() *authProvider {
	return &authProvider{}
}

func (a *authProvider) DBConfig() config.DBConfig {
	a.dbConfig = config.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}
	return a.dbConfig
}

func (a *authProvider) DBClient() *gorm.DB {
dsn := os.Getenv("PG_DSN")

if a.dbClient == nil {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database", err)
	}
	a.dbClient = db
	}
	return a.dbClient
}

func (a  *authProvider) AuthRepo() repository.AuthRepository {
	if a.authRepo == nil {
		a.authRepo = authRepository.NewAuthRepository(a.DBClient())
	}
	return a.authRepo
}

func (a *authProvider) AuthService() service.AuthService {
	if a.authService == nil {
		a.authService = authService.NewAuthService(a.AuthRepo())
	}
	return a.authService
}

func (a *authProvider) AuthImpl() api.AuthApi {
	if a.authImpl == nil {
		a.authImpl = authApi.NewAuthApi(a.AuthService())
	}
	return a.authImpl
}