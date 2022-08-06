package ports

import (
	"github.com/alirezazeynali75/simple-to-do-app/data/repos"
	"gorm.io/gorm"
)

type BaseRepos interface {
	BeginTransaction(tx *gorm.DB) repos.Repos
	CommitTransaction() error
	RollbackTransaction() error
}