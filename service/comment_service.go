package service

import (
	"belajar-rest-api-golang/model/web"
	"context"
	"database/sql"
)

type CommentService interface {
	Create(ctx context.Context, tx *sql.Tx, postId int, userId int, request web.CommentCreateRequest) (web.CommentResponse, error)
}
