package repository

import (
	"context"

	"github.com/adafatya/micro-services/inventory-service/internal/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (p *ProductRepository) Create(ctx context.Context, product models.Product) error {
	return p.DB.WithContext(ctx).Create(&product).Error
}
