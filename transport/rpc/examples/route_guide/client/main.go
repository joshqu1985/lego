package main

import (
	"github.com/joshqu1985/lego/transport/naming"
	pb "github.com/joshqu1985/lego/transport/rpc/examples/route_guide/routeguide"
)

func init() {
	_, _ = naming.Init(naming.Config{
		// Source:    "etcd",
		// Endpoints: []string{"127.0.0.1:9379"},
		Source:    "nacos",
		Endpoints: []string{"127.0.0.1:8848"},
		Cluster:   "f5fc10a5-0f22-4188-ac9c-1e34023f6556",
	})
}

func main() {
	client := NewRecordRouter(naming.Get())

	// Looking for a valid feature
	client.PrintFeature(&pb.Point{Latitude: 409146138, Longitude: -746188906})

	// Feature missing.
	client.PrintFeature(&pb.Point{Latitude: 0, Longitude: 0})

	// Looking for features between 40, -75 and 42, -73.
	client.PrintFeatures(&pb.Rectangle{
		Lo: &pb.Point{Latitude: 400000000, Longitude: -750000000},
		Hi: &pb.Point{Latitude: 420000000, Longitude: -730000000},
	})

	// RecordRoute
	client.RunRecordRoute()

	// RouteChat
	client.RunRouteChat()
}
