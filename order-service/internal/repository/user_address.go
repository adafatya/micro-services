package repository

import (
	"context"

	"github.com/adafatya/micro-services/order-service/internal/models"
	"gorm.io/gorm"
)

type UserAddressRepository struct {
	DB *gorm.DB
}

func NewUserAddressRepository(db *gorm.DB) *UserAddressRepository {
	return &UserAddressRepository{
		DB: db,
	}
}

func (u *UserAddressRepository) AddAddress(ctx context.Context, data *models.UserAddress) (*models.UserAddress, error) {
	err := u.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
