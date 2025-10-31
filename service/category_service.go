package service

import (
	"belajar-rest-api-golang/model/web"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
}
