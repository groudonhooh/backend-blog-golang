package web

import "time"

/*
product response untuk response yang dikirim ke client
sehingga jika ada data yang sensitif tidak perlu dikirim ke client
*/
type PostResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Content   string    `json:"content"`
	ImageURL  string    `json:"image_url"`
	AuthorId  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}
