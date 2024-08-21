package dto

import "mime/multipart"

type AddProductRequest struct {
	ProductName string                  `form:"product_name"`
	Description string                  `form:"description"`
	Price       int                     `form:"price"`
	Quantity    int                     `form:"quantity"`
	ThumbnailID int                     `form:"thumbnail_id"`
	Images      []*multipart.FileHeader `form:"images[]"`
}
