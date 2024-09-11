package api

import (
	"github.com/gin-gonic/gin"
)


type AuthApi interface {
	Login(ctx *gin.Context)
	Signup(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type MovieApi interface {
	AddMovie(ctx *gin.Context)
	UpdateMovie(ctx *gin.Context)
	DeleteMovie(ctx *gin.Context)
	SetSchedules(ctx *gin.Context)
	ShowSchedules(ctx *gin.Context)
	DeleteSchedules(ctx *gin.Context)
}

type SeatsApi interface {
	ReserveSeat(ctx *gin.Context)
	UnreserveSeat(ctx *gin.Context)
	ShowReservedSeats(ctx *gin.Context)
}

