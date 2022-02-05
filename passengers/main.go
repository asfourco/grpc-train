package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	pb "github.com/asfourco/grpc-train/passengers/api/goclient/v1"
)

const (
	listenAddress = "0.0.0.0:9090"
)

type passengersService struct {
	pb.UnimplementedPassengersServiceServer
}

func (p *passengersService) CreatePassenger(ctx context.Context, request *pb.CreatePassengerRequest) (*pb.CreatePassengerResponse, error) {
	return &pb.CreatePassengerResponse{Passenger: &pb.Passenger{
		Id:   uuid.NewString(),
		Name: "Joeseph Bach",
	}}, nil
}

func (p *passengersService) ListPassengers(ctx context.Context, request *pb.ListPassengersRequest) (*pb.ListPassengersResponse, error) {
	return &pb.ListPassengersResponse{Passengers: []*pb.Passenger{
		{
			Id:   uuid.NewString(),
			Name: "Madam Curie",
		},

		{
			Id:   uuid.NewString(),
			Name: "Peter Choi",
		},
	}}, nil
}

func main() {
	log.Printf("Passengers service starting on %s", listenAddress)
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPassengersServiceServer(s, &passengersService{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
