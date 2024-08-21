package services

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/adafatya/micro-services/webapi/internal/dto"
	"github.com/adafatya/micro-services/webapi/internal/repository"
	"github.com/adafatya/micro-services/webapi/pkg/constants"
	"github.com/adafatya/micro-services/webapi/pkg/util"

	pb "github.com/adafatya/micro-services/webapi/internal/proto/inventoryservice"
)

type ProductService struct {
	InventoryServiceClient pb.InventoryServiceClient
	GCSRepository          *repository.GCSRepository
}

func NewProductService(inventoryServiceClient pb.InventoryServiceClient, gcsRepository *repository.GCSRepository) *ProductService {
	return &ProductService{
		InventoryServiceClient: inventoryServiceClient,
		GCSRepository:          gcsRepository,
	}
}

func (p *ProductService) AddProduct(ctx context.Context, data dto.AddProductRequest) (dto.MessageResponse, error) {
	var imagePaths []string
	for i, image := range data.Images {
		file, err := image.Open()
		if err != nil {
			return dto.MessageResponse{Message: fmt.Sprintf("Gagal menambahkan produk: %v", err.Error())}, err
		}

		today := time.Now().Format(time.RFC3339)
		imagePath := fmt.Sprintf("%s%s_%d_%s%s", constants.ImageFolder, data.ProductName, i+1, today, path.Ext(image.Filename))
		if err := p.GCSRepository.Upload(ctx, file, imagePath); err != nil {
			return dto.MessageResponse{Message: fmt.Sprintf("Gagal menambahkan produk: %v", err.Error())}, err
		}

		imagePaths = append(imagePaths, imagePath)
	}

	var productImages []*pb.ProductImage
	for i, imagePath := range imagePaths {
		publicImagePath := fmt.Sprintf("%s%s/%s", constants.GCSPath, util.GetEnv("GCS_BUCKET", ""), imagePath)
		isThumbnail := i+1 == data.ThumbnailID
		productImages = append(productImages, &pb.ProductImage{
			ImagePath:   &publicImagePath,
			IsThumbnail: &isThumbnail,
		})
	}

	price := int32(data.Price)
	quantity := int32(data.Quantity)
	in := &pb.AddProductRequest{
		ProductName:   &data.ProductName,
		Description:   &data.Description,
		Price:         &price,
		Quantity:      &quantity,
		ProductImages: productImages,
	}

	resp, err := p.InventoryServiceClient.AddProduct(ctx, in)
	return dto.MessageResponse{Message: resp.GetMessage()}, err
}
