package repo

import (
	"errors"
	"github.com/alirezazeynali75/simple-to-do-app/data/database/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
	transaction *gorm.DB
}

func (userRepo UserRepo) BeginTransaction(tx *gorm.DB) Repos {
	if tx != nil {
		userRepo.transaction = tx
		return &userRepo
	} else {
		tx := userRepo.db.Begin()
		return &UserRepo{userRepo.db, tx}
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

func (userRepo *UserRepo) getDb() *gorm.DB {
	var db *gorm.DB
	if (userRepo.transaction != nil) {
		db = userRepo.transaction
	} else {
		db = userRepo.db
	}
	return db
}
func (userRepo *UserRepo) FindByPk(id uint) (*model.User, error) {
	dbUser := new(model.User)
	db := userRepo.getDb()
	result := db.First(&dbUser, "id = ?",id)
	if result.Error != nil {
		return nil, result.Error
	}
	return dbUser, nil
}

func (userRepo *UserRepo) Create(user *model.User) (bool, error) {
	db := userRepo.getDb()
	result := db.Create(user)
	_, error := result.Rows()
	if error != nil {
		return false, error
	}
	return true, nil
}

func (userRepo *UserRepo) FindAll() ([]model.User, error) {
	dbUsers := make([]model.User, 0)
	db := userRepo.getDb()
	result := db.Select("name", "family_name", "email", "user_name", "national_id", "is_activated").Find(&dbUsers)
	if (result.Error != nil) {
		return nil, result.Error
	}
	users := make([]model.User, len(dbUsers))
	for _, u := range dbUsers {
		users = append(users, u)
	}
	return users, nil
}

func (userRepo *UserRepo) FindActiveUser() ([]model.User, error) {
	dbUsers := make([]model.User, 0)
	db := userRepo.getDb()
	result := db.Select("name", "family_name", "email", "user_name", "national_id", "is_activated").Where(&model.User{IsActivated: true}).Find(dbUsers)
	if (result.Error != nil) {
		return nil, result.Error
	}
	users := make([]model.User, len(dbUsers))
	for _, u := range dbUsers {
		users = append(users, u)
	}
	return users, nil
}

func (userRepo *UserRepo) FindByUsername(username string) (*model.User, error) {
	dbUser := new(model.User)
	db := userRepo.getDb()
	result := db.First(dbUser, "username = ?", username)
	if (result.Error != nil) {
		return nil, result.Error
	}
	return dbUser, nil
}