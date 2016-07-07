package db

import (
	"errors"
	"log"
	"math"
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

func Paginate(c *gin.Context) (*gorm.DB, error) {
	db := DBInstance(c)
	limit_query := c.DefaultQuery("limit", "25")
	page_query := c.Query("page")
	last_id_query := c.Query("last_id")
	order := c.DefaultQuery("order", "desc")

	limit, err := strconv.Atoi(limit_query)
	if err != nil {
		return db, errors.New("invalid parameter.")
	}
	limit = int(math.Max(1, math.Min(10000, float64(limit))))

	if page_query != "" {
		// pagination 1
		page, err := strconv.Atoi(page_query)
		if err != nil {
			return db, errors.New("invalid parameter.")
		}
		return db.Offset(limit * (page - 1)).Limit(limit), nil
	} else if last_id_query != "" {
		// pagination 2
		last_id, err := strconv.Atoi(last_id_query)
		if err != nil {
			return db, errors.New("invalid parameter.")
		}
		if order == "desc" {
			return db.Where("id < ?", last_id).Limit(limit).Order("id desc"), nil
		} else {
			return db.Where("id > ?", last_id).Limit(limit).Order("id asc"), nil
		}
	}
	return db, nil
}
