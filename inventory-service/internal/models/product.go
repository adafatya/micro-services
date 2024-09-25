package models

type Product struct {
	ID          int
	ProductName string
	Description string
	Price       int
	Quantity    int
	Images      []ProductImage `gorm:"foreignKey:ProductID;references:ID"`
}
