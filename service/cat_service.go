package service

import (
	pb "cloudrun/gen/go/proto/v1"
	"context"
	"errors"
)

type MyCatService struct{}

func (s *MyCatService) GetMyCat(ctx context.Context, message *pb.GetMyCatRequest) (*pb.MyCatResponse, error) {
	message.Name = "tama"
	switch message.Name {
	case "tama":
		return &pb.MyCatResponse{
			Name: "tama",
			Kind: "mainecoon",
		}, nil
	case "mike":
		return &pb.MyCatResponse{
			Name: "mike",
			Kind: "Norwegian Forest Cat",
		}, nil
	}
	return nil, errors.New("not found your cat")
}
