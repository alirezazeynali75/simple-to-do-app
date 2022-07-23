package repo

import (
	"errors"
	"gorm.io/gorm"
	"github.com/alirezazeynali75/simple-to-do-app/model"
)

type UserRepo struct {
	db *gorm.DB
	transaction *gorm.DB
}

func (userRepo *UserRepo) BeginTransaction(tx *gorm.DB) UserRepo {
	if tx != nil {
		userRepo.transaction = tx
		return *userRepo
	} else {
		tx := userRepo.db.Begin()
		return UserRepo{userRepo.db, tx}
	}
}

func (userRepo *UserRepo) CommitTransaction() (bool, error) {
	if userRepo.transaction == nil {
		return false, errors.New("transaction not began yet")
	}
	userRepo.transaction.Commit()
	return true, nil
}

func (userRepo *UserRepo) RollbackTransaction() (bool, error) {
	if userRepo.transaction == nil {
		return false, errors.New("transaction not began yet")
	}
	userRepo.transaction.Rollback()
	return true, nil
}

func (userRepo *UserRepo) findByPk() (User)