package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewSummary struct {
	MovieID       uint    `json:"movie_id"`
	AverageRating float64 `json:"average_rating"`
	ReviewsCount  int     `json:"reviews_count"`
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("[review-service] %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}

func main() {
	r := gin.New()
	r.Use(gin.Recovery(), LoggingMiddleware())

	r.GET("/reviews/:movieId/summary", func(c *gin.Context) {
		movieIDParam := c.Param("movieId")
		movieID, err := strconv.Atoi(movieIDParam)
		if err != nil || movieID <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid movie id"})
			return
		}
		summary := ReviewSummary{
			MovieID:       uint(movieID),
			AverageRating: 8.4,
			ReviewsCount:  127,
		}

		c.JSON(http.StatusOK, summary)
	})

	log.Println("ReviewService running on http://localhost:8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
