package repos

import (
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
