package movieApi

import (
	"github.com/Owariq/go-movie-reserv/internal/api"
	movieModels "github.com/Owariq/go-movie-reserv/internal/models/movies"
	"github.com/Owariq/go-movie-reserv/internal/service"
	"github.com/gin-gonic/gin"
)

type movieApi struct {
	movieService service.MovieService
}

func NewMovieApi(movieService service.MovieService) api.MovieApi {
	return &movieApi{movieService: movieService}
}

	// AddMovie implements api.MovieApi
	// 
	// Add movie
	// 
	//	@Summary		Add movie
	//	@Description	Responds with the movie ID
	//	@Tags			movies
	//	@Accept			json
	//	@Produce		json
	//	@Param			movie	body		movieModels.Movie	true	"Movie info"
	//	@Success		201		{object}	gin.H				"movie added"
	//	@Failure		400		{object}	gin.H				"invalid request"
	//	@Failure		500		{object}	gin.H				"could not add movie"
	//	@Router			/movies [post]
func (m *movieApi) AddMovie(ctx *gin.Context) {
	var movie movieModels.Movie

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	movieID, err := m.movieService.AddMovie(ctx, movie)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, gin.H{"movieID": movieID})
}


	// UpdateMovie implements api.MovieApi
	// 
	// Update movie
	// 
	//	@Summary		Update movie
	//	@Description	Responds with the success message
	//	@Tags			movies
	//	@Accept			json
	//	@Produce		json
	//	@Param			movie	body		movieModels.Movie	true	"Movie info"
	//	@Success		200		{object}	gin.H				"movie updated"
	//	@Failure		400		{object}	gin.H				"invalid request"
	//	@Failure		500		{object}	gin.H				"could not update movie"
	//	@Router			/movies [put]
func (m *movieApi) UpdateMovie(ctx *gin.Context) {
	var movie movieModels.Movie

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := m.movieService.UpdateMovie(ctx, movie)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"success": "movie updated"})
}
	// DeleteMovie implements api.MovieApi
	// 
	// Delete movie
	// 
	//	@Summary		Delete movie
	//	@Description	Responds with the success message
	//	@Tags			movies
	//	@Accept			json
	//	@Produce		json
	//	@Param			movie	body		movieModels.Movie	true	"Movie info"
	//	@Success		200		{object}	gin.H				"movie deleted"
	//	@Failure		400		{object}	gin.H				"invalid request"
	//	@Failure		500		{object}	gin.H				"could not delete movie"
	//	@Router			/movies [delete]
func (m *movieApi) DeleteMovie(ctx *gin.Context) {
	var movieID movieModels.Movie
	
	err := ctx.ShouldBindJSON(&movieID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = m.movieService.DeleteMovie(ctx, movieID.ID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"success": "movie deleted"})
}


	// SetSchedules implements api.MovieApi
	// 
	// Set schedules for a movie
	// 
	//	@Summary		Set schedules for a movie
	//	@Description	Responds with the success message
	//	@Tags			movies
	//	@Accept			json
	//	@Produce		json
	//	@Param			schedules	body		movieModels.Schedules	true	"Schedules info"
	//	@Success		200			{object}	gin.H					"schedules set"
	//	@Failure		400			{object}	gin.H					"invalid request"
	//	@Failure		500			{object}	gin.H					"could not set schedules"
	//	@Router			/movies/schedules [put]
func (m *movieApi) SetSchedules(ctx *gin.Context) {
	var schedules movieModels.Schedules

	if err := ctx.ShouldBindJSON(&schedules); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := m.movieService.SetSchedules(ctx, schedules)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"success": "schedules set"})
}
	// ShowSchedules implements api.MovieApi
	// 
	// Show schedules for a movie
	// 
	//	@Summary		Show schedules for a movie
	//	@Description	Responds with the schedules info
	//	@Tags			movies
	//	@Accept			json
	//	@Produce		json
	//	@Param			movie_id	body		uint	true	"Movie ID"
	//	@Success		200			{object}	gin.H	"schedules"
	//	@Failure		400			{object}	gin.H	"invalid request"
	//	@Failure		500			{object}	gin.H	"could not show schedules"
	//	@Router			/movies/schedules [get]
func (m *movieApi) ShowSchedules(ctx *gin.Context) {

var schedules movieModels.Schedules

if err := ctx.ShouldBindJSON(&schedules); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

schedules, err := m.movieService.ShowSchedules(ctx, schedules.MovieID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"schedules": schedules})
}
	// DeleteSchedules implements api.MovieApi
	// 
	// Delete schedules for a movie
	// 
	//	@Summary		Delete schedules for a movie
	//	@Description	Responds with the success message
	//	@Tags			movies
	//	@Accept			json
	//	@Produce		json
	//	@Param			movie_id	body		uint	true	"Movie ID"
	//	@Success		200			{object}	gin.H	"schedules deleted"
	//	@Failure		400			{object}	gin.H	"invalid request"
	//	@Failure		500			{object}	gin.H	"could not delete schedules"
	//	@Router			/movies/schedules [delete]
func (m *movieApi) DeleteSchedules(ctx *gin.Context) {

	var schedules movieModels.Schedules

	if err := ctx.ShouldBindJSON(&schedules); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := m.movieService.DeleteSchedules(ctx, schedules.MovieID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"success": "schedules deleted"})
}