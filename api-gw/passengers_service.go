package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/asfourco/grpc-train/api-gw/api/goclient/v1"
	passengersSvc "github.com/asfourco/grpc-train/passengers/api/goclient/v1"
)

type passengersService struct {
	passengersClient passengersSvc.PassengersServiceClient
}

// nolint
func NewPassengersService(passengersClient passengersSvc.PassengersServiceClient) *passengersService {
	return &passengersService{
		passengersClient: passengersClient,
	}
}

func (p *passengersService) ListPassengers(ctx context.Context, request *pb.ListPassengersRequest) (*pb.ListPassengersResponse, error) {
	allPassengers, err := p.passengersClient.ListPassengers(ctx, &passengersSvc.ListPassengersRequest{})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result := &pb.ListPassengersResponse{Passengers: make([]*pb.Passenger, len(allPassengers.GetPassengers()))}
	return result, nil
}

func (p *passengersService) CreatePassenger(ctx context.Context, request *pb.CreatePassengerRequest) (*pb.CreatePassengerResponse, error) {
	res, err := p.passengersClient.CreatePassenger(ctx, &passengersSvc.CreatePassengerRequest{
		Id:   request.GetId(),
		Name: request.GetName(),
	})
	if err != nil {
		return nil, err
	}

	result := &pb.CreatePassengerResponse{Passenger: &pb.Passenger{
		Id:   res.GetPassenger().GetId(),
		Name: res.GetPassenger().GetName(),
	}}

	return result, nil
}
