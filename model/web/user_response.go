package web

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token,omitempty"`
}
