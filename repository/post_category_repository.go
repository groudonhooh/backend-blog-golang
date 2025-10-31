package repository

import (
	"context"
	"database/sql"
)

type PostCategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, postId int, categoryId []int) error
	DeleteByPostId(ctx context.Context, tx *sql.Tx, postId int) error
}
