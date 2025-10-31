package domain

import "time"

type Category struct {
	Id        int
	Name      string
	Slug      string
	CreatedAt time.Time
}
