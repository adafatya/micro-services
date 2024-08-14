package repositories

import (
	"context"

	"github.com/adafatya/micro-services/auth-service/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) Create(ctx context.Context, data *models.User) error {
	return u.DB.WithContext(ctx).Create(data).Error
}
