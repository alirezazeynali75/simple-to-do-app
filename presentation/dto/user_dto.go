package dto

type UserSignUpDto struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginDto struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}