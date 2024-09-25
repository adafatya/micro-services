package services

import (
	"context"
	"fmt"

	"github.com/adafatya/micro-services/order-service/internal/models"
	"github.com/adafatya/micro-services/order-service/internal/proto/inventoryservice"
	"github.com/adafatya/micro-services/order-service/internal/repository"

	pb "github.com/adafatya/micro-services/order-service/internal/proto/orderservice"
)

type OrderService struct {
	InventoryServiceClient inventoryservice.InventoryServiceClient
	OrderRepository        *repository.OrderRepository
}

func NewOrderService(inventoryServiceClient inventoryservice.InventoryServiceClient, orderRepository *repository.OrderRepository) *OrderService {
	return &OrderService{
		InventoryServiceClient: inventoryServiceClient,
		OrderRepository:        orderRepository,
	}
}

func (o *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.MessageResponse, error) {
	order := models.Order{
		UserID:        int(in.GetUserID()),
		UserAddressID: int(in.GetUserAddressID()),
	}
	var products []*inventoryservice.Product

	for _, product := range in.GetProducts() {
		products = append(products, &inventoryservice.Product{
			ID:       product.ID,
			Quantity: product.Quantity,
		})
		order.Products = append(order.Products, models.OrderProduct{
			ProductID: int(product.GetID()),
			Quantity:  int(product.GetQuantity()),
		})
	}

	buy, err := o.InventoryServiceClient.BuyProducts(ctx, &inventoryservice.BuyProductsRequest{Products: products})
	fmt.Println(err.Error())
	if err != nil {
		return &pb.MessageResponse{Message: buy.Message}, nil
	}
	order.TotalPrice = buy.GetTotalPrice()
	order.ApprovalStatus = 0

	msg := "Berhasil membuat order!"
	err = o.OrderRepository.AddOrder(ctx, order)
	if err != nil {
		msg = fmt.Sprintf("Gagal membuat order: %v", err.Error())
	}

	return &pb.MessageResponse{Message: &msg}, err
}
