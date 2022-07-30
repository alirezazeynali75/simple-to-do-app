package repos

import (
	"github.com/alirezazeynali75/simple-to-do-app/data/database/mysql/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	BaseRepos
}

func (r *UserRepo) BeginTransaction(tx *gorm.DB) Repos {
	if tx != nil {
		r.Transaction = tx
		return r
	} else {
		tx := r.Db.Begin()
		base := BaseRepos{
			Db: r.Db,
			Transaction: tx,
		}
		return &UserRepo{
			BaseRepos: base,
		}
	}
}

func (r *UserRepo) List() ([]model.User, error) {
	dbUsers := make([]model.User, 0)
	db := r.getDb()
	result := db.Select("username", "email").Find(&dbUsers)
	if (result.Error != nil) {
		return nil, result.Error
	}
	users := make([]model.User, len(dbUsers))
	for _, u := range dbUsers {
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepo) Create(username string, password string, email string, phoneNumber string) (bool, error) {
	user := &model.User{
		Username: username,
		Password: password,
		Email: email,
		PhoneNumber: phoneNumber,
	}
	db := r.getDb()
	db.Create(user)
	return true, nil
}

func (r *UserRepo) UpdateById(id uint, username string, password string, email string) (bool, error) {
	db := r.getDb()
	dbu := new(model.User)
	err := db.First(&dbu, id).Error
	dbu.Password = password
	dbu.Email = email
	dbu.Username = username
	if err != nil {
		return false, err
	}
	err = db.Save(&dbu).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *UserRepo) Delete(id uint) error {
	db := r.getDb()
	err := db.Delete(&model.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}