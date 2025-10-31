package domain

import "time"

type Post struct {
	Id         int
	Title      string
	Slug       string
	Content    string
	ImageURL   string
	AuthorId   int
	CategoryId []int
	CreatedAt  time.Time
}
