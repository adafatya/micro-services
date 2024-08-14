package config

import (
	"github.com/adafatya/micro-services/webapi/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/adafatya/micro-services/webapi/internal/proto/authservice"
)

func NewAuthServiceClient() pb.AuthServiceClient {
	authServiceAddr := util.GetEnv("AUTH_SERVICE_ADDR", "")
	conn, err := grpc.NewClient(authServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	c := pb.NewAuthServiceClient(conn)

	return c
}
