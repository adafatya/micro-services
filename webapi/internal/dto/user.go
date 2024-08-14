package dto

type UserRegisterRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

type UserRegisterResponse struct {
	Message string `json:"message"`
}
