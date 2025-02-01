package rpcserver

import (
	"google.golang.org/grpc"

  "{{$.ServerName}}/internal/service"
	pb "{{$.ServerName}}/rpcserver/protocols"
)

// New 创建RPCServer
func New(svc *service.Service) *RPCServer {
	return &RPCServer{
		service: svc,
	}
}

// RPCServer
type RPCServer struct {
	service *service.Service
}

func (this *RPCServer) Routers(s *grpc.Server) {
	pb.Register{{$.Name}}Server(s, New{{$.Name}}Api(this.service))
}
