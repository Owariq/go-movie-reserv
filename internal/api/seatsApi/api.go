package seatsApi

import (
	"github.com/Owariq/go-movie-reserv/internal/api"
	seatsModel "github.com/Owariq/go-movie-reserv/internal/models/seats"
	"github.com/Owariq/go-movie-reserv/internal/service"
	"github.com/gin-gonic/gin"
)

type seatsApi struct {
	serv service.SeatsService
}

func NewSeatsApi(serv service.SeatsService) api.SeatsApi {
	return &seatsApi{serv: serv}
}

// ReserveSeat implements api.SeatsApi
//
// Reserve a seat
//
//	@Summary		Reserve a seat
//	@Description	Responds with the reserved seat
//	@Tags			seats
//	@Accept			json
//	@Produce		json
//	@Param			seat	body		seatsModel.Seats	true	"Seat info"
//	@Success		200		{object}	seatsModel.Seats	"seat reserved"
//	@Failure		400		{object}	gin.H				"invalid request"
//	@Failure		500		{object}	gin.H				"could not reserve seat"
//	@Router			/seats [post]
func (s *seatsApi) ReserveSeat(ctx *gin.Context) {
	var seat seatsModel.Seats
	if err := ctx.ShouldBindJSON(&seat); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := s.serv.ReserveSeat(ctx, seat)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, seat)
}

// UnreserveSeat implements api.SeatsApi
//
// Unreserve a seat
//
//	@Summary		Unreserve a seat
//	@Description	Responds with the success message
//	@Tags			seats
//	@Accept			json
//	@Produce		json
//	@Param			seat	body		seatsModel.Seats	true	"Seat info"
//	@Success		200		{object}	gin.H				"seat unreserved"
//	@Failure		400		{object}	gin.H				"invalid request"
//	@Failure		500		{object}	gin.H				"could not unreserve seat"
//	@Router			/seats [delete]
func (s *seatsApi) UnreserveSeat(ctx *gin.Context) {
	var seat seatsModel.Seats
	if err := ctx.ShouldBindJSON(&seat); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := s.serv.UnreserveSeat(ctx, seat)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"success": "seat unreserved"})
}

// ShowReservedSeats implements api.SeatsApi
//
// Show reserved seats for a movie
//
//	@Summary		Show reserved seats for a movie
//	@Description	Responds with the reserved seats info
//	@Tags			seats
//	@Accept			json
//	@Produce		json
//	@Param			seat	body		seatsModel.Seats	true	"Seat info"
//	@Success		200		{object}	seatsModel.Seats	"reserved seats"
//	@Failure		400		{object}	gin.H				"invalid request"
//	@Failure		500		{object}	gin.H				"could not show reserved seats"
//	@Router			/seats/reserved [get]
func (s *seatsApi) ShowReservedSeats(ctx *gin.Context) {
	var seat seatsModel.Seats
	if err := ctx.ShouldBindJSON(&seat); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	seats, err := s.serv.ShowReservedSeats(ctx, seat.MovieID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, seats)
}

