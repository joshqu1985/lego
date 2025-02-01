package main

import (
	"github.com/joshqu1985/lego/transport/rest"
)

func RouterRegister(r *rest.Router) {
	h := NewApiHandler()
	r.GET("/hello", rest.ResponseWrapper(h.Hello))
}
