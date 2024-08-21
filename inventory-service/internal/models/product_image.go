package models

type ProductImage struct {
	ID          int
	ProductID   int
	ImagePath   string
	IsThumbnail bool
}
