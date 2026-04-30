package clients

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type ReviewSummary struct {
	MovieID       uint    `json:"movie_id"`
	AverageRating float64 `json:"average_rating"`
	ReviewsCount  int     `json:"reviews_count"`
}

func GetReviewSummary(movieID uint) (*ReviewSummary, error) {
	baseURL := os.Getenv("REVIEW_SERVICE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8081"
	}

	client := resty.New().
		SetBaseURL(baseURL).
		SetTimeout(5 * time.Second)

	client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		log.Printf("[resty][request] %s %s", req.Method, req.URL)
		return nil
	})

	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		log.Printf("[resty][response] status=%d path=%s", resp.StatusCode(), resp.Request.URL)
		return nil
	})

	var summary ReviewSummary
	_, err := client.R().
		SetHeader("Accept", "application/json").
		SetResult(&summary).
		Get(fmt.Sprintf("/reviews/%d/summary", movieID))
	if err != nil {
		return nil, err
	}

	return &summary, nil
}
