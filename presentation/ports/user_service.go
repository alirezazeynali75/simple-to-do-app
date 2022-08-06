package ports

import "github.com/alirezazeynali75/simple-to-do-app/core/interfaces"

type UserService interface {
	Login(email string, password string) (string, *interfaces.User, error)
	UpdatePassword(id uint, password string) error
	SignUp(username string, email string, phoneNumber string, password string) (*interfaces.User, error)
	List() ([]interfaces.User, error)
}