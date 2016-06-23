package db

import (
	"log"

	"github.com/shimastripe/gouserapi/middleware"
	"github.com/shimastripe/gouserapi/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DB  *gorm.DB
	err error
)

func Connect(c *gin.Engine) {
	DB, err = gorm.Open("sqlite3", "db/users.db")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	DB.AutoMigrate(&models.User{})
	c.Use(middleware.SetDBtoContext(DB))
}

func DBInstance(c *gin.Context) *gorm.DB {
	return c.MustGet("DB").(*gorm.DB)
}
