package services

import (
	"context"

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
