package domain

import (
	"time"
	"github.com/alirezazeynali75/simple-to-do-app/utils"
)

type User struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	FamilyName string `json:"family_name"`
	Password string `json:"-"`
	Email string `json:"email"`
	Username string `json:"username"`
	IsActivated bool `json:"is_activated"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-"`
}

func (u *User) SetPassword(password string) error {
	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}

func (u *User) CheckPassword(password string) bool {
	isVerified, err := utils.VerifyHash(password, u.Password)
	if err != nil {
		return false
	}
	return isVerified
}

func (u *User) CreateNewUser(name string, familyName string, email string, password string, username string) (*User, error) {
	user := &User{
		Name: name,
		FamilyName: name,
		Email: email,
		Username: username,
	}
	err := user.SetPassword(password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

type UserService interface {
	SignUp() (bool, error)
	Login() (string, User, error)
	List() ([]User, error)
	UpdatePassword() (bool, error)
}