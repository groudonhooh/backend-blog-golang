package service

import (
	"belajar-rest-api-golang/exception"
	"belajar-rest-api-golang/helper"
	"belajar-rest-api-golang/model/domain"
	"belajar-rest-api-golang/model/web"
	"belajar-rest-api-golang/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
	PostRepository    repository.PostRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewCommentService(commentRepository repository.CommentRepository, postRepository repository.PostRepository, DB *sql.DB, validate *validator.Validate) CommentService {
	return &CommentServiceImpl{
		CommentRepository: commentRepository,
		PostRepository:    postRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *CommentServiceImpl) Create(ctx context.Context, tx *sql.Tx, postId int, userId int, request web.CommentCreateRequest) (web.CommentResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err = service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.PostRepository.FindById(ctx, tx, postId)
	if err != nil {
		panic(exception.NotFoundError{Error: "Post is not found"})
	}

	comment := domain.Comment{
		Id:       0,
		PostId:   postId,
		AuthorId: userId,
		Content:  request.Content,
	}

	comment = service.CommentRepository.Save(ctx, tx, comment)

	return helper.ToCommentResponse(comment), nil
}
