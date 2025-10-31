package web

type PostUpdateRequest struct {
	Id       int    `json:"id"`
	Title    string `json:"title" validate:"required"`
	Content  string `json:"content" validate:"required"`
	ImageURL string `json:"image_url"`
}
