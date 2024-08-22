package services

import (
	"context"
	"fmt"

	"github.com/adafatya/micro-services/order-service/internal/models"
	"github.com/adafatya/micro-services/order-service/internal/repository"

	pb "github.com/adafatya/micro-services/order-service/internal/proto/orderservice"
)

type UserAddressService struct {
	UserAddressRepository *repository.UserAddressRepository
}

func NewUserAddressService(userAddressRepository *repository.UserAddressRepository) *UserAddressService {
	return &UserAddressService{
		UserAddressRepository: userAddressRepository,
	}
}

func (u *UserAddressService) AddAddress(ctx context.Context, data *pb.AddUserAddressRequest) (*pb.AddUserAddressResponse, error) {
	userAddress := models.UserAddress{
		UserID:    int(data.GetUserID()),
		Alamat:    data.GetAlamat(),
		KodePos:   data.GetKodePos(),
		Kelurahan: data.GetKelurahan(),
		Kecamatan: data.GetKecamatan(),
		Kabupaten: data.GetKabupaten(),
		Provinsi:  data.GetProvinsi(),
	}

	msg := "Berhasil menambah alamat!"
	_, err := u.UserAddressRepository.AddAddress(ctx, &userAddress)
	if err != nil {
		msg = fmt.Sprintf("Gagal menambah data: %v", err.Error())
		return &pb.AddUserAddressResponse{Message: &msg}, err
	}

	id := int32(userAddress.ID)
	alamat := userAddress.Alamat + ", " + userAddress.Kelurahan + ", " + userAddress.Kecamatan + ", " + userAddress.Kabupaten + ", " + userAddress.Provinsi + ", " + userAddress.KodePos
	return &pb.AddUserAddressResponse{Message: &msg, ID: &id, AlamatLengkap: &alamat}, nil
}
