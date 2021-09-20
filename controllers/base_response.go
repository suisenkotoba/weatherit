package controller

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Meta    struct {
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	// Pagination interface{} `json:"pagination,omitempty"`
}

func NewSuccessResponse(c echo.Context, param interface{}) error {
	response := BaseResponse{}
	response.Message = "Success"
	response.Data = param
	// response.Pagination = pagination

	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	// response.Message = message
	response.Meta.Messages = []string{err.Error()}

	return c.JSON(status, response)
}
