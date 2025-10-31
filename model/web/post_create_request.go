package web

type PostCreateRequest struct {
	Title      string `json:"title" validate:"required"`
	Slug       string `json:"slug"`
	Content    string `json:"content" validate:"required"`
	ImageURL   string `json:"image_url"`
	CategoryId []int  `json:"category_id" validate:"required"`
}
