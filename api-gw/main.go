package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/asfourco/grpc-train/api-gw/api/goclient/v1"
	passengersSvc "github.com/asfourco/grpc-train/passengers/api/goclient/v1"
	trainsSvc "github.com/asfourco/grpc-train/trains/api/goclient/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	listenAddress        = "0.0.0.0:9090"
	passengersSvcAddress = "passengers:9090"
	trainsSvcAddress     = "trains:9090"
)

func newPassengersSvcClient() (passengersSvc.PassengersServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), passengersSvcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("trains client: %w", err)
	}
	return passengersSvc.NewPassengersServiceClient(conn), nil
}

func newTrainsSvcClient() (trainsSvc.TrainsServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), trainsSvcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("passengers client: %w", err)
	}
	return trainsSvc.NewTrainsServiceClient(conn), nil
}

func logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("method %q called\n", info.FullMethod)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("method %q failed: %s\n", info.FullMethod, err)
	}
	return resp, err
}

func main() {
	log.Printf("GRPC Train API service starting on %s", listenAddress)

	trainsClient, err := newTrainsSvcClient()
	if err != nil {
		panic(err)
	}

	passengersClient, err := newPassengersSvcClient()
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(logger))

	//nolint
	pb.RegisterPassengersServiceServer(s, NewPassengersService(passengersClient))
	//nolint
	pb.RegisterTrainsServiceServer(s, NewTrainsService(trainsClient))

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
