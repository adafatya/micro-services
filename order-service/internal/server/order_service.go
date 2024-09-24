package server

import (
	"context"

	"github.com/adafatya/micro-services/order-service/internal/services"

	pb "github.com/adafatya/micro-services/order-service/internal/proto/orderservice"
)

type OrderServiceServer struct {
	pb.UnimplementedOrderServiceServer

	UserAddressService *services.UserAddressService
	OrderService       *services.OrderService
}

func NewOrderServiceServer(userAddressService *services.UserAddressService, orderService *services.OrderService) *OrderServiceServer {
	return &OrderServiceServer{
		UserAddressService: userAddressService,
		OrderService:       orderService,
	}
}

func (o *OrderServiceServer) AddUserAddress(ctx context.Context, in *pb.AddUserAddressRequest) (*pb.AddUserAddressResponse, error) {
	return o.UserAddressService.AddAddress(ctx, in)
}

func (o *OrderServiceServer) GetUserAddresses(ctx context.Context, in *pb.GetUserAddressesRequest) (*pb.GetUserAddressesResponse, error) {
	return o.UserAddressService.GetUserAddresses(ctx, in)
}

func (o *OrderServiceServer) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.MessageResponse, error) {
	return o.OrderService.CreateOrder(ctx, in)
}
