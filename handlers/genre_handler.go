package handlers

import (
	"backend/config"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetGenres(c *gin.Context) {
	var genres []models.Genre
	config.DB.Order("id asc").Find(&genres)
	c.JSON(http.StatusOK, genres)
}

func CreateGenre(c *gin.Context) {
	var genre models.Genre
	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if genre.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Genre name is required"})
		return
	}
	if err := config.DB.Create(&genre).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create genre"})
		return
	}
	c.JSON(http.StatusCreated, genre)
}

func GetGenreByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var genre models.Genre
	if err := config.DB.First(&genre, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
		return
	}
	c.JSON(http.StatusOK, genre)
}

func DeleteGenre(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := config.DB.Delete(&models.Genre{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete genre"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Genre deleted successfully"})
}
