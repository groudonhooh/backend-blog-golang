package controller

import (
	"belajar-rest-api-golang/helper"
	"belajar-rest-api-golang/model/web"
	"belajar-rest-api-golang/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UserControllerImpl struct {
	UserService service.UserService
	Validate    *validator.Validate
}

func NewUserController(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
		Validate:    validator.New(),
	}
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request) {
	var registerRequest web.UserRegisterRequest
	helper.ReadFromRequestBody(request, &registerRequest)

	err := controller.Validate.Struct(registerRequest)
	helper.PanicIfError(err)

	response, err := controller.UserService.Register(request.Context(), nil, registerRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request) {
	var loginRequest web.UserLoginRequest
	helper.ReadFromRequestBody(request, &loginRequest)

	err := controller.Validate.Struct(loginRequest)
	helper.PanicIfError(err)

	token, err := controller.UserService.Login(request.Context(), nil, loginRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: map[string]string{
			"token": token,
		},
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Logout(writer http.ResponseWriter, request *http.Request) {
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Logged Out",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
