package repo

import "gorm.io/gorm"

type Repos interface {
	BeginTransaction(tx *gorm.DB) Repos
	CommitTransaction() (bool, error)
	RollbackTransaction() (bool, error)
}