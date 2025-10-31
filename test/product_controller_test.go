package test

// import (
// 	"belajar-rest-api-golang/app"
// 	"belajar-rest-api-golang/controller"
// 	"belajar-rest-api-golang/helper"
// 	"belajar-rest-api-golang/middleware"
// 	"belajar-rest-api-golang/repository"
// 	"belajar-rest-api-golang/service"
// 	"database/sql"
// 	"net/http"
// 	"testing"
// 	"time"

// 	"github.com/go-playground/validator/v10"
// )

// func setupTestDB() *sql.DB {
// 	db, err := sql.Open("mysql", "root:alsya12345@tcp(localhost:3306)/golang_restful_api_test")
// 	helper.PanicIfError(err)

// 	db.SetConnMaxLifetime(time.Minute * 60)
// 	db.SetMaxOpenConns(20)
// 	db.SetMaxIdleConns(5)
// 	db.SetConnMaxIdleTime(time.Minute * 10)

// 	return db
// }

// func setupRouter() http.Handler {
// 	db := setupTestDB()
// 	validate := validator.New()
// 	postRepository := repository.NewPostRepository()
// 	postService := service.NewPostService(postRepository, db, validate)
// 	postController := controller.NewPostController(postService)
// 	userRepository := repository.NewUserRepository()
// 	userService := service.NewUserService(userRepository, db, validate)
// 	userController := controller.NewUserController(userService)

// 	router := app.NewRouter(postController, userController)
// 	return middleware.NewAuthMiddleware(router)
// }

// func TestCreateProductSucces(t *testing.T) {

// }

// func TestCreateProductFailed(t *testing.T) {

// }

// func TestUpdateProductSuccess(t *testing.T) {

// }

// func TestUpdateProductFailed(t *testing.T) {

// }

// func TestDeleteProductSuccess(t *testing.T) {

// }

// func TestDeleteProductFailed(t *testing.T) {

// }

// func TestGetProductSuccess(t *testing.T) {

// }

// func TestGetProductFailed(t *testing.T) {

// }

// func TestLIstProductSuccess(t *testing.T) {

// }

// func TestUnauthorized(t *testing.T) {

// }
