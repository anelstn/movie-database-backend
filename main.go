package main

import (
	"backend/config"
	"backend/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	r := gin.Default()

	r.GET("/movies", handlers.GetMovies)
	r.POST("/movies", handlers.CreateMovie)
	r.GET("/movies/:id", handlers.GetMovieByID)
	r.PUT("/movies/:id", handlers.UpdateMovie)
	r.DELETE("/movies/:id", handlers.DeleteMovie)

	r.GET("/directors", handlers.GetDirectors)
	r.POST("/directors", handlers.CreateDirector)
	r.GET("/directors/:id", handlers.GetDirectorByID)
	r.DELETE("/directors/:id", handlers.DeleteDirector)

	r.GET("/genres", handlers.GetGenres)
	r.POST("/genres", handlers.CreateGenre)
	r.GET("/genres/:id", handlers.GetGenreByID)
	r.DELETE("/genres/:id", handlers.DeleteGenre)

	log.Println("Movie Database API is running on http://localhost:8080")
	r.Run(":8080")
}
