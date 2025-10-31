package service

import (
	"belajar-rest-api-golang/model/web"
	"context"
	"database/sql"
)

type UserService interface {
	Login(ctx context.Context, tx *sql.Tx, request web.UserLoginRequest) (string, error)
	Register(ctx context.Context, tx *sql.Tx, request web.UserRegisterRequest) (web.UserResponse, error)
}
