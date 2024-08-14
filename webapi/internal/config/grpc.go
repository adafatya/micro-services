package config

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/adafatya/micro-services/webapi/internal/proto/authservice"
)

func NewAuthServiceClient() pb.AuthServiceClient {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	c := pb.NewAuthServiceClient(conn)

	return c
}
