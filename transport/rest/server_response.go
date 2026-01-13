package rest

import (
	"github.com/gin-gonic/gin"
)

type (
	// RouterHandler gin router回调函数 handler函数定义格式.
	RouterHandler func(ctx *gin.Context) *JSONResponse

	// JSONResponse 返回结构.
	JSONResponse struct {
		Data any    `json:"data,omitempty"`
		Msg  string `json:"msg,omitempty"`
		Code int    `json:"code"`
	}
)

// ResponseWrapper gin router调用封装.
func ResponseWrapper(handle RouterHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := handle(ctx)
		ctx.JSON(response.Code/1000, response)
	}
}

// ErrorResponse 错误返回.
func ErrorResponse(code int, msg string) *JSONResponse {
	return &JSONResponse{Code: code, Msg: msg}
}

// SuccessResponse 正确返回.
func SuccessResponse(data any) *JSONResponse {
	return &JSONResponse{Code: 200000, Msg: "", Data: data}
}
