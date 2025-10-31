package service

import (
	"belajar-rest-api-golang/exception"
	"belajar-rest-api-golang/helper"
	"belajar-rest-api-golang/middleware"
	"belajar-rest-api-golang/model/domain"
	"belajar-rest-api-golang/model/web"
	"belajar-rest-api-golang/repository"
	"context"
	"database/sql"
	"strings"

	"github.com/go-playground/validator/v10"
)

type PostServiceImpl struct {
	PostRepository         repository.PostRepository
	PostCategoryRepository repository.PostCategoryRepository
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewPostService(postRepository repository.PostRepository, postCategoryRepository repository.PostCategoryRepository, DB *sql.DB, validate *validator.Validate) PostService {
	return &PostServiceImpl{
		PostRepository:         postRepository,
		PostCategoryRepository: postCategoryRepository,
		DB:                     DB,
		Validate:               validate,
	}
}

func generatePostSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}

func (service *PostServiceImpl) Create(ctx context.Context, request web.PostCreateRequest, userId int) web.PostResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post := domain.Post{
		Id:         0,
		Title:      request.Title,
		Slug:       generatePostSlug(request.Title),
		Content:    request.Content,
		ImageURL:   request.ImageURL,
		AuthorId:   userId,
		CategoryId: request.CategoryId,
	}

	post = service.PostRepository.Save(ctx, tx, post)

	if len(request.CategoryId) > 0 {
		err := service.PostCategoryRepository.Create(ctx, tx, post.Id, request.CategoryId)
		helper.PanicIfError(err)
	}

	return helper.ToPostResponse(post)
}
func (service *PostServiceImpl) Update(ctx context.Context, request web.PostUpdateRequest) web.PostResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post, err := service.PostRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Ambil user ID dari context
	userId := middleware.GetUserIDFromContext(ctx)
	if userId == 0 {
		panic(exception.NewUnauthorizedError("User ID not found in context"))
	}

	// Cek apakah user adalah pemilik post
	if post.AuthorId != userId {
		panic(exception.NewUnauthorizedError("You are not allowed to update this post"))
	}

	post.Title = request.Title
	post.Content = request.Content
	post.ImageURL = request.ImageURL

	post = service.PostRepository.Update(ctx, tx, post)

	return helper.ToPostResponse(post)
}

func (service *PostServiceImpl) Delete(ctx context.Context, postId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post, err := service.PostRepository.FindById(ctx, tx, postId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	userId := middleware.GetUserIDFromContext(ctx)
	if userId == 0 {
		panic(exception.NewUnauthorizedError("User ID not found in context"))
	}

	// Cek apakah user adalah pemilik post
	if post.AuthorId != userId {
		panic(exception.NewUnauthorizedError("You are not allowed to delete this post"))
	}

	service.PostRepository.Delete(ctx, tx, post)
}

func (service *PostServiceImpl) FindById(ctx context.Context, postId int) web.PostResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post, err := service.PostRepository.FindById(ctx, tx, postId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPostResponse(post)
}

func (service *PostServiceImpl) FindAll(ctx context.Context) []web.PostResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	posts := service.PostRepository.FindAll(ctx, tx)

	return helper.ToPostResponses(posts)
}
