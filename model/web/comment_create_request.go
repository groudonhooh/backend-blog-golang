package web

type CommentCreateRequest struct {
	Content string `json:"content" validate:"required"`
}
