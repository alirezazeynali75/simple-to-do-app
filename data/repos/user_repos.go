package repos

import (
	"errors"

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

func (r *UserRepo) Create(username string, password string, email string, phoneNumber string) (*model.User, error) {
	user := &model.User{
		Username: username,
		Password: password,
		Email: email,
		PhoneNumber: phoneNumber,
	}
	db := r.getDb()
	db.Create(user)
	return user, nil
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

func (r *UserRepo) GetByEmail(email string) (*model.User, error) {
	db := r.getDb()
	user := new(model.User) 
	result := db.First(user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepo) IsExistByEmail(email string) (bool, error) {
	user, err := r.GetByEmail(email)
	if err != nil {
		return false, nil
	}
	if user.Email != email {
		return false, errors.New("query result is not ok")
	}
	return true, nil
}

func (r *UserRepo) GetByUsername(username string) (*model.User, error) {
	db := r.getDb()
	var user *model.User = &model.User{}
	result := db.Model(&model.User{}).First(user, &model.User{Username: username})
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepo) IsExistByUsername(username string) (bool, error) {
	user, err := r.GetByUsername(username)
	if err != nil {
		return false, nil
	}
	if user.Username != username {
		return false, errors.New("query result is not ok")
	}
	return true, nil
}