package services

import (
	"context"

	"github.com/adafatya/micro-services/webapi/internal/dto"

	pb "github.com/adafatya/micro-services/webapi/internal/proto/orderservice"
)

type UserAddressService struct {
	OrderServiceClient pb.OrderServiceClient
}

func NewUserAddressService(orderServiceClient pb.OrderServiceClient) *UserAddressService {
	return &UserAddressService{
		OrderServiceClient: orderServiceClient,
	}
}

func (u *UserAddressService) AddUserAddress(ctx context.Context, data dto.AddUserAddressRequest) (dto.AddUserAddressResponse, error) {
	userID := int32(data.UserID)
	userAddress := &pb.AddUserAddressRequest{
		UserID:    &userID,
		Alamat:    &data.Alamat,
		Kelurahan: &data.Kelurahan,
		Kecamatan: &data.Kecamatan,
		Kabupaten: &data.Kabupaten,
		Provinsi:  &data.Provinsi,
		KodePos:   &data.KodePos,
	}

	resp, err := u.OrderServiceClient.AddUserAddress(ctx, userAddress)
	return dto.AddUserAddressResponse{Message: resp.GetMessage(), ID: int(resp.GetID()), AlamatLengkap: resp.GetAlamatLengkap()}, err
}

func (u *UserAddressService) GetUserAddresses(ctx context.Context, userID int32) (dto.GetUserAddressesResponse, error) {
	resp, err := u.OrderServiceClient.GetUserAddresses(ctx, &pb.GetUserAddressesRequest{UserID: &userID})
	if err != nil {
		return dto.GetUserAddressesResponse{Message: resp.GetMessage()}, err
	}

	var userAddresses []dto.UserAddress
	for _, userAddress := range resp.UserAddresses {
		userAddresses = append(userAddresses, dto.UserAddress{
			ID:            int(userAddress.GetID()),
			AlamatLengkap: userAddress.GetAlamatLengkap(),
		})
	}

	return dto.GetUserAddressesResponse{Message: resp.GetMessage(), UserAddresses: userAddresses}, nil
}
