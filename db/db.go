package db

import (
	"log"
	"path/filepath"
	"strconv"

	"github.com/shimastripe/gouserapi/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Connect() *gorm.DB {
	dir := filepath.Dir("db/database.db")
	db, err := gorm.Open("sqlite3", dir+"/database.db")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	db.AutoMigrate(&models.User{}, &models.Profile{}, &models.AccountName{}, &models.Email{}, &models.Nation{})
	return db
}

func DBInstance(c *gin.Context) *gorm.DB {
	return c.MustGet("DB").(*gorm.DB)
}

func Paginate(c *gin.Context) *gorm.DB {
	db := DBInstance(c)
	limit_query := c.DefaultQuery("limit", "25")
	page_query := c.Query("page")
	last_id_query := c.Query("last_id")
	order := c.Query("order")

	limit, err := strconv.Atoi(limit_query)
	if err != nil {
		limit = 25
	}

	if page_query != "" {
		// pagination 1
		page, err := strconv.Atoi(page_query)
		if err != nil {
			page = 1
		}
		offset := limit * (page - 1)
		return db.Offset(offset).Limit(limit)
	} else if last_id_query != "" {
		// pagination 2

	}

	return db
}
