package repository

import (
	"belajar-rest-api-golang/model/domain"
	"context"
	"database/sql"
)

type CommentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, comment domain.Comment) domain.Comment
}
