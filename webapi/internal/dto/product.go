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

type GetProductsRequest struct {
	Limit   int32  `form:"limit"`
	Page    int32  `form:"page"`
	Keyword string `form:"keyword"`
	Order   string `form:"order"`
}

func (p *GetProductsRequest) GetOffset() int32 {
	return p.Limit * (p.Page - 1)
}

type ProductResponse struct {
	ID            int32  `json:"id"`
	ProductName   string `json:"product_name"`
	Price         int32  `json:"price"`
	ThumbnailPath string `json:"thumbnail_path"`
}

type GetProductsResponse struct {
	Message  string            `json:"message"`
	Products []ProductResponse `json:"products"`
}

type CreateOrderProductRequest struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}
