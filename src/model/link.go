package model

import "time"

type Link struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" validate:"required,min=3"`
	URL       string    `json:"url" validate:"required,url"`
	CreatedAt time.Time `json:"created_at"`
}