package services

import (
	"errors"

	"github.com/alirezazeynali75/simple-to-do-app/core/interfaces"
	"github.com/alirezazeynali75/simple-to-do-app/core/ports"
	"github.com/alirezazeynali75/simple-to-do-app/utils"
)

type UserService struct {
	Repo ports.UserRepo
	JwtManager interfaces.JwtManager
}

func (us *UserService) SignUp(username string, email string, phoneNumber string, password string) (*interfaces.User, error) {
	var isExist bool
	var err error
	isExist, err = us.Repo.IsExistByEmail(email)
	isExist, err = us.Repo.IsExistByUsername(username)
	if isExist {
		return nil, errors.New("user is exist with this email or username")
	}
	passwordHash, err := utils.HashPassword(password)
	createdUser, err:= us.Repo.Create(
		username,
		passwordHash,
		email,
		phoneNumber,
	)
	if err != nil {
		return nil, err
	}
	user := interfaces.CreateUser(createdUser)
	return user, nil
}

func (us *UserService) List() ([]interfaces.User, error) {
	dbUsers, err := us.Repo.List()
	if err != nil {
		return nil, err
	}
	users := []interfaces.User{}
	for _, value := range dbUsers {
		users = append(users, *interfaces.CreateUser(&value))
	}
	return users, err
}

func (us *UserService) Login(email string, password string) (string, *interfaces.User, error) {
	user, err := us.Repo.GetByEmail(email)
	isVerified, err := utils.VerifyHash(password, user.Password)
	if !isVerified {
		return "", nil, errors.New("password which entered is not correct")
	}
	domainUser := interfaces.CreateUser(user)
	bearerToken, err := us.JwtManager.Generate(domainUser)
	if err != nil {
		return "", nil, err
	}
	return bearerToken, domainUser, nil
}

func (us *UserService) UpdatePassword(id uint, password string) error {
	return errors.New("can not update password")
}


