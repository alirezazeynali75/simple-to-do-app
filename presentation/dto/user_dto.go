package dto

type UserSignUpDto struct {
	Username string `json:"username"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password string `json:"password"`
}