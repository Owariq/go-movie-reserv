package authService

import (
	"context"
	"fmt"
	"os"
	"time"

	models "github.com/Owariq/go-movie-reserv/internal/models/auth"
	"github.com/Owariq/go-movie-reserv/internal/repository"
	"github.com/Owariq/go-movie-reserv/internal/service"
	utils "github.com/Owariq/go-movie-reserv/internal/utils/jwt"
	"github.com/dgrijalva/jwt-go"
)

type serv struct {
	authRepo repository.AuthRepository
}

func NewAuthService(authRepo repository.AuthRepository) service.AuthService {
	return &serv{authRepo: authRepo}
}

func (s *serv) Login(ctx context.Context, email, password string) (string, int, error) {
	user, err := s.authRepo.GetUser(ctx, email)
	if err != nil {
		return "",0,err
	}
	errHash := utils.CompareHashPassword(password, user.Password)

	if !errHash {
		return "",0, fmt.Errorf("invalid password")
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &models.Claims{
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "",0, fmt.Errorf("could not generate token")
	}

	return tokenString,int(expirationTime.Unix()), nil


}

func (s *serv) Signup(ctx context.Context, user models.User) error {
	
	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)
	 if errHash != nil {
        return fmt.Errorf("could not generate password hash")
    }

	err := s.authRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}