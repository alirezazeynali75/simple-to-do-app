package interfaces

import (
	"time"

	"github.com/alirezazeynali75/simple-to-do-app/data/database/mysql/model"
)

type User struct {
	Id uint `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password string		`json:"-"`
	IsActivated bool `json:"is_activated"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
	DelatedAt time.Time	`json:"-"`
}

func CreateUser(domainModel *model.User) *User {
	user := &User{
		Id: domainModel.ID,
		Email: domainModel.Email,
		PhoneNumber: domainModel.PhoneNumber,
		Password: domainModel.Password,
		IsActivated: domainModel.IsActivated,
		CreatedAt: domainModel.CreatedAt,
		UpdateAt: domainModel.UpdatedAt,
		DelatedAt: domainModel.DeletedAt.Time,
	}
	return user
}

type UserService interface {
	Login(email string, password string) (string, *User, error)
	UpdatePassword(id uint, password string) error
	SignUp(username string, email string, phoneNumber string, password string) (*User, error)
	List() ([]User, error)
}