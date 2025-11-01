package web

type PostCreateRequest struct {
	Title       string `json:"title" validate:"required"`
	Slug        string `json:"slug"`
	Content     string `json:"content" validate:"required"`
	ImageURL    string `json:"image_url"`
	CategoryIds []int  `json:"category_ids" validate:"required"`
}
