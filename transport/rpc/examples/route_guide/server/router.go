package main

import (
	"google.golang.org/grpc"

	pb "github.com/joshqu1985/lego/transport/rpc/examples/route_guide/routeguide"
)

func RouterRegister(s *grpc.Server) {
	pb.RegisterRouteGuideServer(s, NewRouteGuide())
}
