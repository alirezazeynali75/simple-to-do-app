package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/alirezazeynali75/simple-to-do-app/utils"
)

func OpenConnection() *gorm.DB {
	config := utils.Config{}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: config.DbDriver,
		DSN: config.GetDbSource(),
	}))
	if err != nil {
		panic(err)
	}
	return db
}

func KillConnection(db *gorm.DB) error {
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	sqlDb.Close()
	return nil
}
