package server

import (
	"context"
	"fmt"

	"github.com/adafatya/micro-services/inventory-service/internal/models"
	pb "github.com/adafatya/micro-services/inventory-service/internal/proto/inventoryservice"
	"github.com/adafatya/micro-services/inventory-service/internal/services"
)

type InventoryServiceServer struct {
	pb.UnimplementedInventoryServiceServer

	ProductService *services.ProductService
}

func NewInventoryServiceServer(productService *services.ProductService) *InventoryServiceServer {
	return &InventoryServiceServer{
		ProductService: productService,
	}
}

func (s *InventoryServiceServer) AddProduct(ctx context.Context, in *pb.AddProductRequest) (*pb.MessageResponse, error) {
	data := models.Product{
		ProductName: in.GetProductName(),
		Description: in.GetDescription(),
		Price:       int(in.GetPrice()),
		Quantity:    int(in.GetQuantity()),
	}

	for _, img := range in.GetProductImages() {
		data.Images = append(data.Images, models.ProductImage{
			ImagePath:   img.GetImagePath(),
			IsThumbnail: img.GetIsThumbnail(),
		})
	}

	msg := "Berhasil menambah produk!"
	err := s.ProductService.AddProduct(ctx, data)
	if err != nil {
		msg = fmt.Sprintf("Gagal menambah produk: %v", err.Error())
	}

	return &pb.MessageResponse{Message: &msg}, err
}

func (s *InventoryServiceServer) GetProducts(ctx context.Context, in *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	msg := "Berhasil mendapatkan daftar produk!"

	products, err := s.ProductService.GetProducts(ctx, in)
	if err != nil {
		msg = fmt.Sprintf("Gagal mendapatkan produk: %v", err.Error())
	}

	return &pb.GetProductsResponse{Message: &msg, Products: products}, err
}

func (s *InventoryServiceServer) BuyProducts(ctx context.Context, in *pb.BuyProductsRequest) (*pb.BuyProductsResponse, error) {
	return s.ProductService.BuyProducts(ctx, in)
}
