package services

import (
	"context"

	"github.com/adafatya/micro-services/inventory-service/internal/models"
	"github.com/adafatya/micro-services/inventory-service/internal/repository"
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
