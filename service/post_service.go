package service

import (
	"belajar-rest-api-golang/model/web"
	"context"
)

/*
membuat kontrak dalam bentuk interface terlebih dahulu
di service ini berupa logic bisnis dari aplikasinya
*/
type PostService interface {
	Create(ctx context.Context, request web.PostCreateRequest, userId int) web.PostResponse
	Update(ctx context.Context, request web.PostUpdateRequest) web.PostResponse
	Delete(ctx context.Context, postId int)
	FindById(ctx context.Context, postId int) web.PostResponse
	FindAll(ctx context.Context) []web.PostResponse
}
