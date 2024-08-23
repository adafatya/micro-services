package dto

type CreateOrderRequest struct {
	UserID        int
	UserAddressID int                         `json:"user_address_id"`
	Products      []CreateOrderProductRequest `json:"products"`
}
