package repository

import (
	"context"

	"github.com/adafatya/micro-services/order-service/internal/dto"
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

func (u *UserAddressRepository) FindUserAddresses(ctx context.Context, userID int32) ([]dto.UserAddressResponse, error) {
	var userAddresses []dto.UserAddressResponse

	address := "CONCAT(alamat, ', ', kelurahan, ', ', kecamatan, ', ', kabupaten, ', ', provinsi, ', ', kode_pos) AS alamat_lengkap"
	if err := u.DB.WithContext(ctx).Model(models.UserAddress{}).Select("id, " + address).Find(&userAddresses).Error; err != nil {
		return nil, err
	}

	return userAddresses, nil
}
