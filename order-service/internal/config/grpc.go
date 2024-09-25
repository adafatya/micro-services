package config

import (
	"github.com/adafatya/micro-services/order-service/internal/proto/inventoryservice"
	"github.com/adafatya/micro-services/order-service/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCServer() *grpc.Server {
	s := grpc.NewServer()
	return s
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
