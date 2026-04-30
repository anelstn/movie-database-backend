package main

import (
	"backend/config"
	"backend/handlers"
	"backend/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	r.POST("/auth/register", handlers.Register)
	r.POST("/auth/login", handlers.Login)

	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/movies", handlers.GetMovies)
		protected.POST("/movies", handlers.CreateMovie)
		protected.GET("/movies/:id", handlers.GetMovieByID)
		protected.GET("/movies/:id/details", handlers.GetMovieWithReviewSummary)
		protected.PUT("/movies/:id", handlers.UpdateMovie)
		protected.DELETE("/movies/:id", handlers.DeleteMovie)

		protected.GET("/directors", handlers.GetDirectors)
		protected.POST("/directors", handlers.CreateDirector)
		protected.GET("/directors/:id", handlers.GetDirectorByID)
		protected.DELETE("/directors/:id", handlers.DeleteDirector)

		protected.GET("/genres", handlers.GetGenres)
		protected.POST("/genres", handlers.CreateGenre)
		protected.GET("/genres/:id", handlers.GetGenreByID)
		protected.DELETE("/genres/:id", handlers.DeleteGenre)
	}

	log.Println("Movie Database API with JWT Authentication")
	log.Println("Server: http://localhost:8080")
	log.Println("POST /auth/register")
	log.Println("POST /auth/login")
	log.Println("Protected routes require Bearer token")
	r.Run(":8080")
}
