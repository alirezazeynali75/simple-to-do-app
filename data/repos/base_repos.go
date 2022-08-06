package repos

import (
	"errors"

	"gorm.io/gorm"
)

type Repos interface {
	BeginTransaction(tx *gorm.DB) Repos
	CommitTransaction() error
	RollbackTransaction() error
}

type BaseRepos struct {
	Db *gorm.DB
	Transaction *gorm.DB
}

func (r *BaseRepos) CommitTransaction() error {
	if r.Transaction == nil {
		return errors.New("transaction not began yet")
	}
	r.Transaction.Commit()
	return nil
}

func (r *BaseRepos) RollbackTransaction() error {
	if r.Transaction == nil {
		return errors.New("transaction not began yet")
	}
	r.Transaction.Rollback()
	return nil
}

func (r *BaseRepos) getDb() *gorm.DB {
	var db *gorm.DB
	if (r.Transaction != nil) {
		db = r.Transaction
	} else {
		db = r.Db
	}
	return db
}



