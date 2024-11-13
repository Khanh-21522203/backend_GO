package gateway

import (
	pb "GoFeed/internal/handlers/grpc_api"
	"context"
)

type GrpcGateway interface {
	CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	GetUser(ctx context.Context, r *pb.GetUserRequest) (*pb.GetUserResponse, error)
	UpdateUser(ctx context.Context, r *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
}

type gateway struct {
	GrpcGateway
	client pb.GoFeedServiceClient
}

func NewGrpcGetway(client pb.GoFeedServiceClient) *gateway {
	return &gateway{
		client: client,
	}
}

func (g *gateway) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return g.client.CreateUser(ctx, &pb.CreateUserRequest{
		UserName:  r.UserName,
		Password:  r.Password,
		FirstName: r.FirstName,
		LastName:  r.LastName,
	})
}

func (g *gateway) GetUser(ctx context.Context, r *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return g.client.GetUser(ctx, &pb.GetUserRequest{
		UserId: r.UserId,
	})
}

func (g *gateway) UpdateUser(ctx context.Context, r *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return g.client.UpdateUser(ctx, &pb.UpdateUserRequest{
		UserId:    r.UserId,
		FirstName: r.FirstName,
		LastName:  r.LastName,
	})
}
