package repository

import (
	"belajar-rest-api-golang/helper"
	"context"
	"database/sql"
)

type PostCategoryRepositoryImpl struct{}

func NewPostCategoryRepository() PostCategoryRepository {
	return &PostCategoryRepositoryImpl{}
}

func (repository *PostCategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, postId int, categoryId []int) error {
	SQL := "INSERT INTO post_categories(post_id, category_id) VALUES (?, ?)"
	for _, categoryId := range categoryId {
		_, err := tx.ExecContext(ctx, SQL, postId, categoryId)
		helper.PanicIfError(err)
	}

	return nil
}

func (repository *PostCategoryRepositoryImpl) DeleteByPostId(ctx context.Context, tx *sql.Tx, postId int) error {
	SQL := "DELETE FROM post_categories WHERE post_id = ?"
	_, err := tx.ExecContext(ctx, SQL, postId)
	helper.PanicIfError(err)

	return nil
}
