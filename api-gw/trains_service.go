package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/asfourco/grpc-train/api-gw/api/goclient/v1"
	trainsSvc "github.com/asfourco/grpc-train/trains/api/goclient/v1"
)

type trainsService struct {
	trainsClient trainsSvc.TrainsServiceClient
}

//nolint
func NewTrainsService(trainsClient trainsSvc.TrainsServiceClient) *trainsService {
	return &trainsService{trainsClient: trainsClient}
}

func (t *trainsService) CreateTrain(ctx context.Context, request *pb.CreateTrainRequest) (*pb.CreateTrainResponse, error) {
	res, err := t.trainsClient.CreateTrain(ctx, &trainsSvc.CreateTrainRequest{
		Id:        request.GetId(),
		From:      request.GetFrom(),
		To:        request.GetFrom(),
		Capacity:  request.GetCapacity(),
		Departure: request.GetDeparture(),
		Arrival:   request.GetArrival(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateTrainResponse{Train: &pb.Train{
		Id:        res.GetTrain().GetId(),
		From:      res.GetTrain().GetFrom(),
		To:        res.GetTrain().GetFrom(),
		Capacity:  res.GetTrain().GetCapacity(),
		Departure: res.GetTrain().GetDeparture(),
		Arrival:   res.GetTrain().GetArrival(),
	}}, nil
}

func (t *trainsService) ListTrains(ctx context.Context, request *pb.ListTrainsRequest) (*pb.ListTrainsResponse, error) {
	allTrains, err := t.trainsClient.ListTrains(ctx, &trainsSvc.ListTrainsRequest{})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result := &pb.ListTrainsResponse{Trains: make([]*pb.Train, len(allTrains.GetTrains()))}
	return result, nil
}
