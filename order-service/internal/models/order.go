package models

type Order struct {
	ID             int
	UserID         int
	UserAddressID  int
	TotalPrice     int64
	ApprovalStatus int
	Products       []OrderProduct `gom:"foreignKey:OrderID;references:ID"`
}
