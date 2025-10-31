package repository

import (
	"belajar-rest-api-golang/model/domain"
	"context"
	"database/sql"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
}
