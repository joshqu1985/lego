package main

import (
	"context"
	"io"
	"log"
	rand "math/rand/v2"
	"time"

	"github.com/joshqu1985/lego/transport/naming"
	"github.com/joshqu1985/lego/transport/rpc"
	pb "github.com/joshqu1985/lego/transport/rpc/examples/route_guide/routeguide"
)

func NewRecordRouter(n naming.Naming) *RecordRouter {
	c, err := rpc.NewClient("route_guide", rpc.WithNaming(n))
	if err != nil {
		log.Fatalf("rpc.NewClient failed: %v", err)
		return nil
	}

	return &RecordRouter{
		client: pb.NewRouteGuideClient(c.Conn()),
	}
}

type RecordRouter struct {
	client pb.RouteGuideClient
}

func (this *RecordRouter) PrintFeature(point *pb.Point) {
	log.Printf("Getting feature for point (%d, %d)", point.Latitude, point.Longitude)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	feature, err := this.client.GetFeature(ctx, point)
	if err != nil {
		log.Fatalf("client.GetFeature failed: %v", err)
	}
	log.Println(feature)
}

func (this *RecordRouter) PrintFeatures(rect *pb.Rectangle) {
	log.Printf("Looking for features within %v", rect)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := this.client.ListFeatures(ctx, rect)
	if err != nil {
		log.Fatalf("client.ListFeatures failed: %v", err)
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.ListFeatures failed: %v", err)
		}
		log.Printf("Feature: name: %q, point:(%v, %v)", feature.GetName(),
			feature.GetLocation().GetLatitude(), feature.GetLocation().GetLongitude())
	}
}

func (this *RecordRouter) RunRecordRoute() {
	// Create a random number of random points
	pointCount := int(rand.Int32N(100)) + 2 // Traverse at least two points
	var points []*pb.Point
	for i := 0; i < pointCount; i++ {
		points = append(points, this.RandomPoint())
	}
	log.Printf("Traversing %d points.", len(points))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := this.client.RecordRoute(ctx)
	if err != nil {
		log.Fatalf("client.RecordRoute failed: %v", err)
	}
	for _, point := range points {
		if err := stream.Send(point); err != nil {
			log.Fatalf("client.RecordRoute: stream.Send(%v) failed: %v", point, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("client.RecordRoute failed: %v", err)
	}
	log.Printf("Route summary: %v", reply)
}

func (this *RecordRouter) RunRouteChat() {
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

	stream, err := this.client.RouteChat(ctx)
	if err != nil {
		log.Fatalf("client.RouteChat failed: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("client.RouteChat failed: %v", err)
			}
			log.Printf("Got message %s at point(%d, %d)", in.Message, in.Location.Latitude, in.Location.Longitude)
		}
	}()
	for _, note := range notes {
		if err := stream.Send(note); err != nil {
			log.Fatalf("client.RouteChat: stream.Send(%v) failed: %v", note, err)
		}
	}
	stream.CloseSend()
	<-waitc
}

func (this *RecordRouter) RandomPoint() *pb.Point {
	lat := (rand.Int32N(180) - 90) * 1e7
	long := (rand.Int32N(360) - 180) * 1e7
	return &pb.Point{Latitude: lat, Longitude: long}
}
