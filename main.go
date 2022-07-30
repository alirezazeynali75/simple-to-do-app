package main

import (
	"github.com/alirezazeynali75/simple-to-do-app/data/database/mysql"
	"github.com/alirezazeynali75/simple-to-do-app/data/database/mysql/model"
	"github.com/alirezazeynali75/simple-to-do-app/utils"
)

func main() {
	utils.LoadEnv()
	db := mysql.OpenConnection()
	defer mysql.KillConnection(db)
	err := db.AutoMigrate(
		&model.Status{},
		&model.Tasks{},
		&model.User{},
	)
	if err != nil {
		panic(err)
	}
}