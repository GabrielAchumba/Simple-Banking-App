package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type appHandler func(ctx *gin.Context) *Response

func ServeHTTP(handle appHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result := handle(ctx) // Controller Action

		if result == nil {
			ctx.JSON(http.StatusInternalServerError, Response{
				StatusCode: http.StatusInternalServerError,
				Message:    "Internal Server Error",
				Data:       nil,
				Success:    false,
			})
		} else {
			ctx.JSON(result.StatusCode, result)
		}
	}
}
