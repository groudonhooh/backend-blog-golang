package web

import "time"

type CommentResponse struct {
	Id        int       `json:"id"`
	PostId    int       `json:"post_id"`
	AuthorId  int       `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
