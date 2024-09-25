package services

import (
	"context"
	"fmt"

	"github.com/adafatya/micro-services/inventory-service/internal/models"
	"github.com/adafatya/micro-services/inventory-service/internal/repository"

	pb "github.com/adafatya/micro-services/inventory-service/internal/proto/inventoryservice"
)

type ProductService struct {
	ProductRepository *repository.ProductRepository
}

func NewProductService(productRepository *repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

func (p *ProductService) AddProduct(ctx context.Context, data models.Product) error {
	return p.ProductRepository.Create(ctx, data)
}

func (p *ProductService) GetProducts(ctx context.Context, query *pb.GetProductsRequest) ([]*pb.Product, error) {
	products, err := p.ProductRepository.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	var resp []*pb.Product
	for _, product := range products {
		id := int32(product.ID)
		price := int32(product.Price)

		productResponse := &pb.Product{
			ID:          &id,
			ProductName: &product.ProductName,
			Price:       &price,
		}

		if len(product.Images) != 0 {
			productResponse.ThumbnailPath = &product.Images[0].ImagePath
		}
		resp = append(resp, productResponse)
	}

	return resp, nil
}

func (p *ProductService) BuyProducts(ctx context.Context, data *pb.BuyProductsRequest) (*pb.BuyProductsResponse, error) {
	var products []models.Product
	for _, product := range data.GetProducts() {
		products = append(products, models.Product{
			ID:       int(product.GetID()),
			Quantity: int(product.GetQuantity()),
		})
	}

	totalPrice, err := p.ProductRepository.Buy(ctx, products)
	msg := "Berhasil mendaftarkan produk pada order!"
	if err != nil {
		msg = fmt.Sprintf("Gagal mendaftarkan produk pada order: %v", err.Error())
		fmt.Println(&msg)
		return &pb.BuyProductsResponse{Message: &msg, TotalPrice: &totalPrice}, err
	}

	return &pb.BuyProductsResponse{Message: &msg, TotalPrice: &totalPrice}, nil
}
