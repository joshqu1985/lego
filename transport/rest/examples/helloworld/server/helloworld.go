package main

import (
	"sync/atomic"

	"github.com/joshqu1985/lego/logs"
	"github.com/joshqu1985/lego/transport/rest"
)

type Handler struct{}

var count int64

func NewApiHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Hello(ctx *rest.Context) *rest.JSONResponse {
	var args struct {
		Data string `binding:"required" form:"data"`
	}

	if err := ctx.ShouldBindQuery(&args); err != nil {
		return rest.ErrorResponse(400000, err.Error())
	}

	atomic.AddInt64(&count, 1)
	if atomic.LoadInt64(&count)%10000 == 0 {
		logs.Info("-------->", atomic.LoadInt64(&count)/10000, "万")
	}
	if atomic.LoadInt64(&count) > 500000 && atomic.LoadInt64(&count) < 2000000 {
		return rest.ErrorResponse(503000, "server busy")
	}

	return rest.SuccessResponse("hello" + args.Data)
}
