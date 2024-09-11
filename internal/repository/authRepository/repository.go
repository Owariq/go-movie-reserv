package authRepository

import (
	"context"
	"fmt"

	models "github.com/Owariq/go-movie-reserv/internal/models/auth"
	"github.com/Owariq/go-movie-reserv/internal/repository"
	"gorm.io/gorm"
)


type repo struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) repository.AuthRepository {
	return &repo{db: db}
}


func (r *repo) GetUser(ctx context.Context, email string) (models.User, error) {
	var user models.User
	r.db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return user, fmt.Errorf("user does not exist")
	}
	return user, nil
}

func (r *repo) CreateUser(ctx context.Context, user models.User) error {
	var existingUser models.User
	
	r.db.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		return fmt.Errorf("user already exists")
	}
	r.db.Create(&user)
	return nil
}