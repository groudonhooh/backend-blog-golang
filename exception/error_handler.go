package exception

import (
	"belajar-rest-api-golang/helper"
	"belajar-rest-api-golang/model/web"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	if unauthorizedError(writer, request, err) {
		return
	}

	if badRequestError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		var message string
		for _, e := range exception {
			switch e.Field() {
			case "Username":
				message = "Username cannot be empty"
			case "Email":
				message = "Email cannot be empty"
			case "Password":
				message = "Password cannot be empty"
			case "Name":
				message = "Name cannot be empty"
			default:
				message = fmt.Sprintf("%s cannot be empty", e.Field())
			}
			break
		}

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   message,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	}
	return false
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not found",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}

}

func unauthorizedError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(UnauthorizedError)
	if ok {
		webResponse := web.WebResponse{
			Code:   http.StatusForbidden, //403
			Status: "Forbidden",
			Data:   exception.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return true
	}
	return false
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func badRequestError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	ex, ok := err.(*ErrorLogin)
	if ok {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   ex.Message,
		}
		helper.WriteToResponseBody(writer, webResponse)
		return true
	}
	return false
}
