package main

import (
	"context"
	"errors"
	"io"
	"log"
	"time"

	rand "math/rand/v2"

	"github.com/joshqu1985/lego/transport/naming"
	"github.com/joshqu1985/lego/transport/rpc"
	pb "github.com/joshqu1985/lego/transport/rpc/examples/route_guide/routeguide"
)

type RecordRouter struct {
	client pb.RouteGuideClient
}

var (
	StringListFeatures = "client.ListFeatures failed:%v"
	StringRecordRoute  = "client.RecordRoute failed:%v"
	StringRouteChat    = "client.RouteChat failed: %v"
)

func NewRecordRouter(n naming.Naming) *RecordRouter {
	c, err := rpc.NewClient("route_guide", rpc.WithNaming(n))
	if err != nil {
		log.Printf("rpc.NewClient failed: %v", err)

		return nil
	}

	return &RecordRouter{
		client: pb.NewRouteGuideClient(c.Conn()),
	}
}

func (rr *RecordRouter) PrintFeature(point *pb.Point) {
	log.Printf("Getting feature for point (%d, %d)", point.GetLatitude(), point.GetLongitude())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	feature, err := rr.client.GetFeature(ctx, point)
	if err != nil {
		log.Printf("client.GetFeature failed: %v", err)

		return
	}
	log.Println(feature)
}

func (rr *RecordRouter) PrintFeatures(rect *pb.Rectangle) {
	log.Printf("Looking for features within %v", rect)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := rr.client.ListFeatures(ctx, rect)
	if err != nil {
		log.Printf(StringListFeatures, err)

		return
	}
	for {
		feature, xerr := stream.Recv()
		if errors.Is(xerr, io.EOF) {
			break
		}
		if xerr != nil {
			log.Printf(StringListFeatures, xerr)

			return
		}
		log.Printf("Feature: name: %q, point:(%v, %v)", feature.GetName(),
			feature.GetLocation().GetLatitude(), feature.GetLocation().GetLongitude())
	}
}

func (rr *RecordRouter) RunRecordRoute() {
	// Create a random number of random points
	pointCount := int(rand.Int32N(100)) + 2 //nolint:gosec // Traverse at least two points
	var points []*pb.Point
	for range pointCount {
		points = append(points, rr.RandomPoint())
	}
	log.Printf("Traversing %d points.", len(points))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := rr.client.RecordRoute(ctx)
	if err != nil {
		log.Printf(StringRecordRoute, err)

		return
	}
	for _, point := range points {
		if xerr := stream.Send(point); xerr != nil {
			log.Printf("client.RecordRoute: stream.Send(%v) failed: %v", point, xerr)

			return
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf(StringRecordRoute, err)

		return
	}
	log.Printf("Route summary: %v", reply)
}

func (rr *RecordRouter) RunRouteChat() {
	notes := []*pb.RouteNote{
		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "First message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Second message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Third message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "Fourth message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Fifth message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Sixth message"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := rr.client.RouteChat(ctx)
	if err != nil {
		log.Printf(StringRouteChat, err)

		return
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, xerr := stream.Recv()
			if errors.Is(xerr, io.EOF) {
				// read done.
				close(waitc)

				return
			}
			if xerr != nil {
				log.Printf(StringRouteChat, xerr)

				return
			}
			log.Printf(
				"Got message %s at point(%d, %d)",
				in.GetMessage(),
				in.GetLocation().GetLatitude(),
				in.GetLocation().GetLongitude(),
			)
		}
	}()
	for _, note := range notes {
		if xerr := stream.Send(note); xerr != nil {
			log.Printf("client.RouteChat: stream.Send(%v) failed: %v", note, xerr)

			return
		}
	}
	_ = stream.CloseSend()
	<-waitc
}

func (rr *RecordRouter) RandomPoint() *pb.Point {
	lat := (rand.Int32N(180) - 90) * 1e7  //nolint:gosec
	lng := (rand.Int32N(360) - 180) * 1e7 //nolint:gosec

	return &pb.Point{Latitude: lat, Longitude: lng}
}
