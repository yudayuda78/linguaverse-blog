package model

import "time"

type Post struct {
	ID int
	Title string
	Slug string
	Content string
	UserID int
	CreatedAt time.Time
	UpdatedAt time.Time
}