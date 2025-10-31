package app

import (
	"belajar-rest-api-golang/controller"
	"belajar-rest-api-golang/exception"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(postController controller.PostController, userController *controller.UserControllerImpl, commentController controller.CommentController, categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/users/register", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		userController.Register(writer, request)
	})
	router.POST("/api/users/login", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		userController.Login(writer, request)
	})
	router.POST("/api/users/logout", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		userController.Logout(writer, request)
	})

	router.POST("/api/categories", categoryController.Create)
	router.GET("/api/categories", categoryController.FindAll)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.GET("/api/posts", postController.FindAll)
	router.GET("/api/posts/:postId", postController.FindById)
	router.POST("/api/posts", postController.Create)
	router.PUT("/api/posts/:postId", postController.Update)
	router.DELETE("/api/posts/:postId", postController.Delete)

	router.POST("/api/posts/:postId/comments", commentController.Create)

	router.PanicHandler = exception.ErrorHandler

	return router
}
