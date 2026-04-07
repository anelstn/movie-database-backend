package handlers

import (
	"backend/config"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	genreFilter := c.Query("genre")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	var movies []models.Movie
	query := config.DB.Preload("Director").Preload("Genre")

	if genreFilter != "" {
		query = query.Joins("JOIN genres ON genres.id = movies.genre_id").
			Where("genres.name = ?", genreFilter)
	}
	query.Order("movies.id asc").
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&movies)
	c.JSON(http.StatusOK, movies)
}

func CreateMovie(c *gin.Context) {
	var input models.MovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movie := models.Movie{
		Title:      input.Title,
		DirectorID: input.DirectorID,
		GenreID:    input.GenreID,
		Year:       input.Year,
		Rating:     input.Rating,
	}
	if err := config.DB.Create(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create movie"})
		return
	}
	config.DB.Preload("Director").Preload("Genre").First(&movie, movie.ID)
	c.JSON(http.StatusCreated, movie)
}

func GetMovieByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var movie models.Movie
	if err := config.DB.Preload("Director").Preload("Genre").First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func UpdateMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var movie models.Movie
	if err := config.DB.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	var input models.MovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie.Title = input.Title
	movie.DirectorID = input.DirectorID
	movie.GenreID = input.GenreID
	movie.Year = input.Year
	movie.Rating = input.Rating

	if err := config.DB.Save(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
		return
	}
	config.DB.Preload("Director").Preload("Genre").First(&movie, id)
	c.JSON(http.StatusOK, movie)
}

func DeleteMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := config.DB.Delete(&models.Movie{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
