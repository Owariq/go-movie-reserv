package app

import (
	"log"
	"os"

	"github.com/Owariq/go-movie-reserv/internal/middleware"
	"github.com/Owariq/go-movie-reserv/internal/utils/closer"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	authProvider *authProvider
	movieProvider *movieProvider
	seatsProvider *seatsProvider
	httpServer *gin.Engine
}

func NewApp() (*App, error) {
	a := &App{}

	err := a.initDeps()
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) Run() error {
	defer func(){
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runHTTPServer()
}

func (a *App) initDeps() error {
	inits := []func() error{
		a.initConfig,
		a.initHTTPServer,
		a.initAuthProvider,
		a.initMovieProvider,
		a.initSeatsProvider,
	}

	for _, init := range inits {
		if err := init(); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initAuthProvider() error {
	a.authProvider = NewAuthProvider()
	return nil
}
func (a *App) initMovieProvider() error {
	a.movieProvider = NewMovieProvider()
	return nil
}
func (a *App) initSeatsProvider() error {
	a.seatsProvider = NewSeatsProvider()
	return nil
}

func (a *App) initHTTPServer() error {
	a.httpServer = gin.Default()
	return nil
}

func (a *App) runHTTPServer() error {
	log.Printf("starting server on port %s", os.Getenv("PORT"))

	a.authProvider.AuthImpl()
	a.movieProvider.MovieImpl()
	a.seatsProvider.SeatsImpl()

	
	a.httpServer.POST("/user/login", a.authProvider.authImpl.Login)
	a.httpServer.POST("/user/signup", a.authProvider.authImpl.Signup)
	a.httpServer.GET("/user/logout", a.authProvider.authImpl.Logout)

	a.httpServer.POST("/movie", middleware.IsAuthorized(), a.movieProvider.movieImpl.AddMovie)
	a.httpServer.PATCH("/movie", middleware.IsAuthorized(), a.movieProvider.movieImpl.UpdateMovie)
	a.httpServer.DELETE("/movie", middleware.IsAuthorized(), a.movieProvider.movieImpl.DeleteMovie)

	a.httpServer.POST("/movie/schedules",  a.movieProvider.movieImpl.SetSchedules)
	a.httpServer.GET("/movie/schedules",  a.movieProvider.movieImpl.ShowSchedules)
	a.httpServer.DELETE("/movie/schedules", a.movieProvider.movieImpl.DeleteSchedules)

	a.httpServer.POST("/seats", a.seatsProvider.seatsImpl.ReserveSeat)
	a.httpServer.DELETE("/seats", a.seatsProvider.seatsImpl.UnreserveSeat)
	a.httpServer.GET("/seats", a.seatsProvider.seatsImpl.ShowReservedSeats)


	
	a.httpServer.GET("/", middleware.IsAuthorized(), func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "welcome"})
	},)
	a.httpServer.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := a.httpServer.Run(":" + os.Getenv("PORT"))
	if err != nil {
		return err
	}

	return nil
}