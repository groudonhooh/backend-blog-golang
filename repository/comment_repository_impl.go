package repository

import (
	"belajar-rest-api-golang/helper"
	"belajar-rest-api-golang/model/domain"
	"context"
	"database/sql"
)

type CommentRepositoryImpl struct{}

func NewCommentRepository() CommentRepository {
	return &CommentRepositoryImpl{}
}

func (repository CommentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, comment domain.Comment) domain.Comment {
	sql := "INSERT INTO comments(content, post_id, author_id) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, sql, comment.Content, comment.PostId, comment.AuthorId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	comment.Id = int(id)
	return comment
}
