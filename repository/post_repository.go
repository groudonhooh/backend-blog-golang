package repository

import (
	"belajar-rest-api-golang/model/domain"
	"context"
	"database/sql"
)

/*
membuat kontrak dalam bentuk interface terlebih dahulu
susunannya yaitu nama function(context, sql transaction/tidak, model) return value nya
*/

type PostRepository interface {
	Save(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post
	Update(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post
	Delete(ctx context.Context, tx *sql.Tx, post domain.Post)
	FindById(ctx context.Context, tx *sql.Tx, postId int) (domain.Post, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Post
	FindAuthorIdByPostId(ctx context.Context, tx *sql.Tx, postId int) (int, error)
}
