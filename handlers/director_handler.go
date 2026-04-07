package handlers

import (
	"backend/config"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDirectors(c *gin.Context) {
	var directors []models.Director
	config.DB.Order("id asc").Find(&directors)
	c.JSON(http.StatusOK, directors)
}

func CreateDirector(c *gin.Context) {
	var director models.Director
	if err := c.ShouldBindJSON(&director); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if director.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Director name is required"})
		return
	}
	if err := config.DB.Create(&director).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create director"})
		return
	}
	c.JSON(http.StatusCreated, director)
}

func GetDirectorByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var director models.Director
	if err := config.DB.First(&director, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Director not found"})
		return
	}
	c.JSON(http.StatusOK, director)
}

func DeleteDirector(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := config.DB.Delete(&models.Director{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete director"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Director deleted successfully"})
}
