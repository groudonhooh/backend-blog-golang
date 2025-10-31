package domain

import "time"

type Comment struct {
	Id        int
	PostId    int
	Content   string
	AuthorId  int
	CreatedAt time.Time
}
