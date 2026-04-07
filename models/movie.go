package models

import "time"

type Movie struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	Title      string  `gorm:"not null;size:255" json:"title" binding:"required,min=1"`
	DirectorID uint    `gorm:"not null" json:"director_id" binding:"required,gt=0"`
	GenreID    uint    `gorm:"not null" json:"genre_id" binding:"required,gt=0"`
	Year       int     `gorm:"not null" json:"year" binding:"required,gt=1900"`
	Rating     float64 `gorm:"not null" json:"rating" binding:"required,gte=0,lte=10"`

	Director Director `gorm:"foreignKey:DirectorID" json:"director"`
	Genre    Genre    `gorm:"foreignKey:GenreID" json:"genre"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MovieInput struct {
	Title      string  `json:"title" binding:"required,min=1"`
	DirectorID uint    `json:"director_id" binding:"required,gt=0"`
	GenreID    uint    `json:"genre_id" binding:"required,gt=0"`
	Year       int     `json:"year" binding:"required,gt=1900"`
	Rating     float64 `json:"rating" binding:"required,gte=0,lte=10"`
}
