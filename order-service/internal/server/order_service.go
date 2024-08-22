package server

import (
	"context"

	"github.com/adafatya/micro-services/order-service/internal/services"

	pb "github.com/adafatya/micro-services/order-service/internal/proto/orderservice"
)

type OrderServiceServer struct {
	pb.UnimplementedOrderServiceServer

	UserAddressService *services.UserAddressService
}

func NewOrderServiceServer(userAddressService *services.UserAddressService) *OrderServiceServer {
	return &OrderServiceServer{
		UserAddressService: userAddressService,
	}
}

func (o *OrderServiceServer) AddUserAddress(ctx context.Context, in *pb.AddUserAddressRequest) (*pb.AddUserAddressResponse, error) {
	return o.UserAddressService.AddAddress(ctx, in)
}

func (o *OrderServiceServer) GetUserAddresses(ctx context.Context, in *pb.GetUserAddressesRequest) (*pb.GetUserAddressesResponse, error) {
	return o.UserAddressService.GetUserAddresses(ctx, in)
}
