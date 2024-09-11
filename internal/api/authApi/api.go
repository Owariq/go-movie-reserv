package authApi

import (
	"errors"

	"github.com/Owariq/go-movie-reserv/internal/api"
	models "github.com/Owariq/go-movie-reserv/internal/models/auth"
	"github.com/Owariq/go-movie-reserv/internal/service"
	"github.com/gin-gonic/gin"
)

type authApi struct {
	authService service.AuthService
}

func NewAuthApi(authService service.AuthService) api.AuthApi {
	return &authApi{authService: authService}
}


// Login implements api.AuthApi
//
// Login user and return token to client in a cookie
//
//	@Summary		Login user
//	@Description	Responds with the token in a cookie
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.User	true	"User credentials"
//	@Success		200		{object}	gin.H		"user logged in"
//	@Failure		400		{object}	gin.H		"invalid password"	"user does not exist"
//	@Failure		500		{object}	gin.H		"could not generate token"
//	@Router			/login [post]
func (a *authApi) Login(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, expirationTime, err := a.authService.Login(ctx, user.Email, user.Password)

	if err != nil {
		if err.Error() == "invalid password" {
			ctx.JSON(400, gin.H{"error": "invalid password"})
			return
		} else if err.Error() == "user does not exist" {
			ctx.JSON(400, gin.H{"error": "user does not exist"})
			return
		} else if err.Error() == "could not generate token" {
			ctx.JSON(500, gin.H{"error": "could not generate token"})
			return
		}
	}

	ctx.SetCookie("token", token, expirationTime, "/", "localhost", false, true)
	ctx.JSON(200, gin.H{"success": "user logged in"})
}

// Signup implements api.AuthApi
//
// Signup user
//
//	@Summary		Signup user
//	@Description	Responds with the success message
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.User	true	"User credentials"
//	@Success		200		{object}	gin.H		"user created"
//	@Failure		400		{object}	gin.H		"user already exists"
//	@Failure		500		{object}	gin.H		"could not generate password hash"
//	@Router			/signup [post]
func (a *authApi) Signup(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := a.authService.Signup(ctx, user)

	var errorMsg error
	if err != nil {
		switch errors.Is(err, errorMsg) {
		case errorMsg == errors.New("user already exists"):
			ctx.JSON(400, gin.H{"error": "user already exists"})
			return
		case errorMsg == errors.New("could not generate password hash"):
			ctx.JSON(500, gin.H{"error": "could not generate password hash"})
			return
		}
	}
	ctx.JSON(200, gin.H{"success": "user created"})
}

//  Logout implements api.AuthApi
//
//  Logout user and delete cookie
//	@Summary		Logout user
//	@Description	Responds with the success message
//	@Tags			auth
//	@Produce		json
//	@Success		200	{object}	gin.H	"user logged out"
//	@Router			/logout [get]
func (a *authApi) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	ctx.JSON(200, gin.H{"success": "user logged out"})
}
