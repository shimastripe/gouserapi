package main

import (
	"go-user-api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	db, _ := gorm.Open("sqlite3", "db/user.db")
	db.CreateTable(&models.User{})
}
