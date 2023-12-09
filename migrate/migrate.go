package main

import (
	"fmt"
	"go-react-todo-udemy/db"
	"go-react-todo-udemy/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
