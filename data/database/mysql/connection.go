package mysql

import (
	"fmt"

	"github.com/alirezazeynali75/simple-to-do-app/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	config := utils.GetConfig()
	fmt.Println(config.GetDbSource())
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: config.DbDriver,
		DSN: config.GetDbSource(),
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
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
