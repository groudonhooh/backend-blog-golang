package domain

import "time"

type Post struct {
	Id          int
	Title       string
	Slug        string
	Content     string
	ImageURL    string
	AuthorId    int
	Author      string
	CategoryIds []int
	CreatedAt   time.Time
}
