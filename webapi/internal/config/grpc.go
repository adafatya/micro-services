package config

import (
	"github.com/adafatya/micro-services/webapi/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/adafatya/micro-services/webapi/internal/proto/authservice"
	"github.com/adafatya/micro-services/webapi/internal/proto/inventoryservice"
	"github.com/adafatya/micro-services/webapi/internal/proto/orderservice"
)

func NewAuthServiceClient() authservice.AuthServiceClient {
	authServiceAddr := util.GetEnv("AUTH_SERVICE_ADDR", "")
	conn, err := grpc.NewClient(authServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	c := authservice.NewAuthServiceClient(conn)

	return c
}

func NewInventoryServiceClient() inventoryservice.InventoryServiceClient {
	inventoryServiceAddr := util.GetEnv("INVENTORY_SERVICE_ADDR", "")
	conn, err := grpc.NewClient(inventoryServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	c := inventoryservice.NewInventoryServiceClient(conn)

	return c
}

func NewOrderServiceClient() orderservice.OrderServiceClient {
	orderServiceAddr := util.GetEnv("ORDER_SERVICE_ADDR", "")
	conn, err := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	c := orderservice.NewOrderServiceClient(conn)

	return c
}
