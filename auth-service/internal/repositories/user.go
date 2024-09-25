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

func (u *UserRepository) FindOne(ctx context.Context, fields string, query string, args ...any) (*models.User, error) {
	var user *models.User

	if err := u.DB.WithContext(ctx).Select(fields).Where(query, args...).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
