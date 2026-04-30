package handlers

import (
	"backend/clients"
	"backend/config"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieDetailsResponse struct {
	Movie         models.Movie           `json:"movie"`
	ReviewSummary *clients.ReviewSummary `json:"review_summary"`
}

func GetMovieWithReviewSummary(c *gin.Context) {
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

	summary, err := clients.GetReviewSummary(movie.ID)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to fetch review summary from review-service"})
		return
	}

	c.JSON(http.StatusOK, MovieDetailsResponse{
		Movie:         movie,
		ReviewSummary: summary,
	})
}
