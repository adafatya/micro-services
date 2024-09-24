package repository

import (
	"context"

	"github.com/adafatya/micro-services/inventory-service/internal/models"
	"gorm.io/gorm"

	pb "github.com/adafatya/micro-services/inventory-service/internal/proto/inventoryservice"
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

func (p *ProductRepository) Find(ctx context.Context, query *pb.GetProductsRequest) ([]models.Product, error) {
	base := p.DB.WithContext(ctx).Model(models.Product{}).Where("quantity > 0")
	base.Preload("Images", func(db *gorm.DB) *gorm.DB {
		return db.Model(models.ProductImage{}).Order("is_thumbnail DESC").Limit(1)
	})

	if query.GetKeyword() != "" {
		base.Where("product_name LIKE ", "%"+query.GetKeyword()+"%")
	}

	switch query.GetOrder() {
	case "terlama":
		base.Order("id")
	case "termurah":
		base.Order("price")
	case "termahal":
		base.Order("price DESC")
	default:
		base.Order("id DESC")
	}

	var products []models.Product
	if err := base.Limit(int(query.GetLimit())).Offset(int(query.GetOffset())).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductRepository) Buy(ctx context.Context, products []models.Product) (int64, error) {
	var totalPrice int64
	totalPrice = 0

	for i := 0; i < len(products); i++ {
		product := models.Product{
			ID: products[i].ID,
		}

		if err := p.DB.WithContext(ctx).First(&product).Error; err != nil {
			return 0, err
		}

		product.Quantity = product.Quantity - products[i].Quantity
		if err := p.DB.WithContext(ctx).Save(&product).Error; err != nil {
			return 0, err
		}

		totalPrice += int64(products[i].Quantity * product.Price)
	}

	return totalPrice, nil
}
