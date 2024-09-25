package services

import (
	"context"

	"github.com/adafatya/micro-services/webapi/internal/dto"
	pb "github.com/adafatya/micro-services/webapi/internal/proto/orderservice"
)

type OrderService struct {
	OrderServiceClient pb.OrderServiceClient
}

func NewOrderService(orderServiceClient pb.OrderServiceClient) *OrderService {
	return &OrderService{
		OrderServiceClient: orderServiceClient,
	}
}

func (o *OrderService) CreateOrder(ctx context.Context, data dto.CreateOrderRequest) (dto.MessageResponse, error) {
	var products []*pb.Product
	for _, product := range data.Products {
		id := int32(product.ID)
		quantity := int32(product.Quantity)
		products = append(products, &pb.Product{
			ID:       &id,
			Quantity: &quantity,
		})
	}

	userID := int32(data.UserID)
	userAddressID := int32(data.UserAddressID)
	in := &pb.CreateOrderRequest{
		UserID:        &userID,
		UserAddressID: &userAddressID,
		Products:      products,
	}

	resp, err := o.OrderServiceClient.CreateOrder(ctx, in)
	return dto.MessageResponse{Message: resp.GetMessage()}, err
}
