package controller

import (
	"belajar-rest-api-golang/helper"
	"belajar-rest-api-golang/middleware"
	"belajar-rest-api-golang/model/web"
	"belajar-rest-api-golang/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CommentControllerImpl struct {
	commentService service.CommentService
}

func NewCommentController(commentService service.CommentService) CommentController {
	return &CommentControllerImpl{
		commentService: commentService,
	}
}

func (controller *CommentControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := middleware.GetUserIDFromContext(request.Context())
	if userId == 0 {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   "No user ID in context",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	commentCreateRequest := web.CommentCreateRequest{}
	helper.ReadFromRequestBody(request, &commentCreateRequest)

	postId := params.ByName("postId")
	id, err := strconv.Atoi(postId)
	helper.PanicIfError(err)

	commentResponse, err := controller.commentService.Create(request.Context(), nil, id, userId, commentCreateRequest)
	if err != nil {
		helper.WriteErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   commentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
