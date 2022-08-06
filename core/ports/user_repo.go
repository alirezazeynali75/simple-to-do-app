package ports

import "github.com/alirezazeynali75/simple-to-do-app/data/database/mysql/model"

type UserRepo interface {
	BaseRepos
	List() ([]model.User, error)
	Create(username string, password string, email string, phoneNumber string) (*model.User, error)
	UpdateById(id uint, username string, password string, email string) (bool, error)
	Delete(id uint) error
	IsExistByEmail(email string) (bool, error)
	IsExistByUsername(username string) (bool, error)
	GetByEmail(email string) (*model.User, error)
}