package main

import (
	"fmt"
	"sync/atomic"

	"github.com/joshqu1985/lego/transport/rest"
)

type Handler struct {
}

func NewApiHandler() *Handler {
	return &Handler{}
}

var count int64

func (h *Handler) Hello(ctx *rest.Context) *rest.JSONResponse {
	var args struct {
		Data string `form:"data" binding:"required"`
	}

	if err := ctx.ShouldBindQuery(&args); err != nil {
		return rest.ErrorResponse(400000, err.Error())
	}

	atomic.AddInt64(&count, 1)
	if atomic.LoadInt64(&count)%10000 == 0 {
		fmt.Println("-------->", atomic.LoadInt64(&count)/10000, "万")
	}
	if atomic.LoadInt64(&count) > 500000 && atomic.LoadInt64(&count) < 2000000 {
		return rest.ErrorResponse(503000, "server busy")
	}
	return rest.SuccessResponse("hello" + args.Data)
}
