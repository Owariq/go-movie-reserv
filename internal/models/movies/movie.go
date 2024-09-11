package movieModels

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title string `json:"title"`
	Description string `json:"description"`
	Year int `json:"year"`
	Genre string `json:"genre"`
	Poster string `json:"poster"`
}