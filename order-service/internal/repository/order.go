package repository

import (
	"context"

	"github.com/adafatya/micro-services/order-service/internal/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (o *OrderRepository) AddOrder(ctx context.Context, order models.Order) error {
	return o.DB.WithContext(ctx).Create(&order).Error
}
