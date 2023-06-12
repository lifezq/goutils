// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package base

import "net/http"

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

func SuccessResponse(data interface{}) *Response {
	return &Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}
}

func ForbiddenResponse(message string) *Response {
	return &Response{
		Code:    http.StatusForbidden,
		Message: message,
	}
}

func BadResponse(message string) *Response {
	return &Response{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func ServerErrorResponse(message string) *Response {
	return &Response{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func ErrorResponse(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
	}
}
