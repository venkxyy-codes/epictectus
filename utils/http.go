package utils

import (
	errs "epictectus/error"
	"net/http"
)

const (
	InvalidRequest = "invalid request"
)

type ValidationError struct {
	Field   string
	Message string
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

type ErrorContent struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Args    interface{} `json:"args"`
}

type ErrorResponse struct {
	Error ErrorContent `json:"error"`
}

type ErrorListResponse struct {
	Errors []ErrorContent `json:"errors"`
}

func RenderSuccess(data interface{}) SuccessResponse {
	response := SuccessResponse{
		Data: data,
	}
	return response
}

func renderErrorMsg(code string, message string, args interface{}) ErrorResponse {
	errorContent := ErrorContent{
		Code:    code,
		Message: message,
		Args:    args,
	}
	response := ErrorResponse{
		Error: errorContent,
	}
	return response
}

func renderErrorListMsg(errorsList []error) ErrorListResponse {
	var errorContents []ErrorContent
	for _, err := range errorsList {
		errorContent := ErrorContent{
			Code:    err.Error(),
			Message: "Error occurred during user creation",
			Args:    nil,
		}
		errorContents = append(errorContents, errorContent)
	}
	return ErrorListResponse{
		Errors: errorContents,
	}
}

// RenderError returns the HTTP status code and error response
func RenderError(err error, args interface{}, customMessage ...string) (int, ErrorResponse) {
	var httpStatus int
	var code string
	var message string

	switch err {
	case errs.ErrInvalidRequest:
		code = err.Error()
		httpStatus = http.StatusBadRequest
		message = "invalid request"
	default:
		code = err.Error()
		httpStatus = http.StatusInternalServerError
		message = "internal server error"
	}

	if len(customMessage) > 0 && customMessage[0] != "" {
		message = customMessage[0]
	}

	if args == nil || args == "" {
		args = err.Error()
	}

	return httpStatus, renderErrorMsg(code, message, args)
}

func RenderErrorList(errorsList []error) (int, ErrorListResponse) {
	return http.StatusInternalServerError, renderErrorListMsg(errorsList)
}

func RenderValidationErrors(errs []ValidationError) (int, ErrorResponse) {
	var errorContents []ErrorContent
	for _, err := range errs {
		errorContent := ErrorContent{
			Code:    InvalidRequest,
			Message: err.Message,
			Args:    err.Field,
		}
		errorContents = append(errorContents, errorContent)
	}
	return http.StatusBadRequest, ErrorResponse{Error: ErrorContent{
		Code:    InvalidRequest,
		Message: "validation error",
		Args:    errorContents,
	}}
}
