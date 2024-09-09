package httputil

import "github.com/gin-gonic/gin"

func NewError(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, HTTPError{
		Code:    status,
		Message: err.Error(),
	})
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
