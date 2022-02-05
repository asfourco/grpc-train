package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "github.com/asfourco/grpc-train/trains/api/goclient/v1"
)

const (
	listenAddress = "0.0.0.0:9090"
)

type trainsService struct {
	pb.UnimplementedTrainsServiceServer
}

func (t *trainsService) CreateTrain(ctx context.Context, request *pb.CreateTrainRequest) (*pb.CreateTrainResponse, error) {
	return &pb.CreateTrainResponse{Train: &pb.Train{
		Id:        "KL009",
		From:      "London",
		To:        "Berlin",
		Capacity:  150,
		Departure: time.Date(2022, 1, 10, 10, 00, 00, 00, time.Local).Format(time.RFC1123Z),
		Arrival:   time.Date(2022, 1, 10, 13, 00, 00, 00, time.Local).Format(time.RFC1123Z),
	}}, nil
}

func (t *trainsService) ListTrains(ctx context.Context, request *pb.ListTrainsRequest) (*pb.ListTrainsResponse, error) {
	return &pb.ListTrainsResponse{Trains: []*pb.Train{
		{
			Id:        "JW120",
			From:      "Belgrade",
			To:        "Berlin",
			Capacity:  300,
			Departure: time.Date(2022, 1, 25, 11, 00, 00, 00, time.Local).Format(time.RFC1123Z),
			Arrival:   time.Date(2022, 1, 25, 19, 00, 00, 00, time.Local).Format(time.RFC1123Z),
		},
		{
			Id:        "TV9980",
			From:      "Helsinki",
			To:        "Prague",
			Capacity:  200,
			Departure: time.Date(2022, 2, 2, 8, 00, 00, 00, time.Local).Format(time.RFC1123Z),
			Arrival:   time.Date(2022, 2, 3, 13, 00, 00, 00, time.Local).Format(time.RFC1123Z),
		},
	}}, nil
}

func main() {
	log.Printf("Trains service starting on %s", listenAddress)
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTrainsServiceServer(s, &trainsService{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
