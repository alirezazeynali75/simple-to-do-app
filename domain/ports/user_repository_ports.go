package ports

import (
	"github.com/alirezazeynali75/simple-to-do-app/data/database/model"
	"github.com/alirezazeynali75/simple-to-do-app/data/database/repo"
	"gorm.io/gorm"
)

type UserRepository interface {
	BeginTransaction(tx *gorm.DB) repo.Repos
	CommitTransaction() (bool, error)
	RollbackTransaction() (bool, error)
	FindByPk(id uint) (*model.User, error)
	Create(user *model.User) (bool, error)
	FindAll() ([]model.User, error)
	FindActiveUser() ([]model.User, error)
}