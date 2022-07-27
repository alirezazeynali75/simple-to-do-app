package services

import (
	"errors"

	"github.com/alirezazeynali75/simple-to-do-app/data/database/model"
	"github.com/alirezazeynali75/simple-to-do-app/domain"
	"github.com/alirezazeynali75/simple-to-do-app/domain/ports"
)

type UserService struct {
	userRepo ports.UserRepository
	jwt *domain.JwtManager
}


func (us *UserService) SignUp(name string, familyName string, email string, password string, username string, nationalId uint) (bool, error) {
	user, err := domain.CreateNewUser(name, familyName, email, password, username)
	if err != nil {
		return false, err
	}
	userModel := &model.User{
		Name: user.Name,
		FamilyName: user.Name,
		Password: user.Password,
		UserName: user.Username,
		NationalId: nationalId,
		Email: user.Email,
		IsActivated: true,
	}
	isCreated, dbErr := us.userRepo.Create(userModel)
	if dbErr != nil {
		return false, dbErr
	}
	return isCreated, nil
}

func (us *UserService) Login(username string, password string) (string, *domain.User, error) {
	dbUser, err := us.userRepo.FindByUsername(username)
	if err != nil {
		return "", nil, err
	}
	user := &domain.User{
		Id: dbUser.ID,
		Name: dbUser.Name,
		FamilyName: dbUser.FamilyName,
		Password: dbUser.Password,
		Email: dbUser.Email,
		Username: dbUser.UserName,
		IsActivated: dbUser.IsActivated,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
	isVerified := user.CheckPassword(password)
	if !isVerified {
		return "", nil, errors.New("password not verified")
	}
	token, jwtErr := us.jwt.Generate(user)
	if jwtErr != nil {
		return "", nil, errors.New("password not verified")
	}
	return token, user, nil
}

func (us *UserService) List() ([]domain.User, error) {
	dbUsers, dbErr := us.userRepo.FindActiveUser()
	if dbErr != nil {
		return nil, dbErr
	}
	users := []domain.User{}
	for _, v := range dbUsers {
		users = append(users, domain.User{
			Id: v.ID,
			Name: v.Name,
			FamilyName: v.FamilyName,
			Password: v.Password,
			Email: v.Email,
			Username: v.UserName,
			IsActivated: v.IsActivated,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt.Time,
		})
	}
	return users, nil
}

func (us *UserService) UpdatedAt() (bool, error) {
	return false, errors.New("method not implemented")
}

