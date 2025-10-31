package main

import (
	"belajar-rest-api-golang/app"
	"belajar-rest-api-golang/controller"
	"belajar-rest-api-golang/helper"
	"belajar-rest-api-golang/middleware"
	"belajar-rest-api-golang/repository"
	"belajar-rest-api-golang/service"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// ====== INISIALISASI DASAR ======
	db := app.NewDB()
	validate := validator.New()

	// ====== PRODUCT SETUP ======
	postRepository := repository.NewPostRepository()
	postCategoryRepository := repository.NewPostCategoryRepository()
	postService := service.NewPostService(postRepository, postCategoryRepository, db, validate)
	postController := controller.NewPostController(postService)

	// ====== USER SETUP ======
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	comentRepository := repository.NewCommentRepository()
	commentService := service.NewCommentService(comentRepository, postRepository, db, validate)
	commentController := controller.NewCommentController(commentService)

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// ====== ROUTER SETUP ======
	router := app.NewRouter(postController, userController, commentController, categoryController)

	// ====== SERVER SETUP ======
	server := http.Server{
		Addr:    "localhost:5000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("Server running at http://localhost:5000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
