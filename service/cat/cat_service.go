package cat

import (
	pb "cloudrun/gen/go/proto/v1"
	"context"
	"errors"
)

type MyCatService struct{}

func (s *MyCatService) GetMyCat(ctx context.Context, message *pb.GetMyCatRequest) (*pb.MyCatResponse, error) {
	switch name := message.GetName(); name {
	case "tama":
		return &pb.MyCatResponse{
			Name: name,
			Kind: "mainecoon",
		}, nil
	case "mike":
		return &pb.MyCatResponse{
			Name: name,
			Kind: "Norwegian Forest Cat",
		}, nil
	}
	return nil, errors.New("not found your cat")
}
