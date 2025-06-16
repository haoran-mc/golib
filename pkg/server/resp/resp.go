package resp

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func HandleSuccess(ctx echo.Context, data any) error {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Code: http.StatusOK, Message: "success", Data: data}
	return ctx.JSON(http.StatusOK, resp)
}

func HandleError(ctx echo.Context, httpCode int, message string, data any) error {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Code: httpCode, Message: message, Data: data}
	return ctx.JSON(httpCode, resp)
}
