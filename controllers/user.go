package controllers

import (
	"net/http"
	"strings"

	dbpkg "github.com/shimastripe/gouserapi/db"
	"github.com/shimastripe/gouserapi/middleware"
	"github.com/shimastripe/gouserapi/models"

	"github.com/blang/semver"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func setPreload(fields string, db *gorm.DB) ([]string, *gorm.DB) {
	list := strings.Split(fields, ",")
	sel := strings.Split(fields, ",")
	offset := 0
	for key, val := range list {
		switch val {
		// Belongs-to
		case "profile":
			db = db.Preload("Profile")
			db = db.Preload("Profile.Nation")
			sel = append(sel[:(key-offset)], sel[(key+1-offset):]...)
			offset += 1
			idflag := true
			for _, v := range sel {
				if v == "profile_id" {
					idflag = false
					break
				}
			}
			if idflag {
				sel = append(sel, "profile_id")
			}
		// Has-one
		case "account_name":
			db = db.Preload("AccountName")
			sel = append(sel[:(key-offset)], sel[(key+1-offset):]...)
			offset += 1
		// Has-many
		case "emails":
			db = db.Preload("Emails")
			sel = append(sel[:(key-offset)], sel[(key+1-offset):]...)
			offset += 1
		case "*":
			db = db.Preload("Profile").Preload("Profile.Nation").Preload("AccountName").Preload("Emails")
		}
	}
	return sel, db
}

func GetUsers(c *gin.Context) {
	version, err1 := middleware.VersionInit(c)
	pagination := dbpkg.Pagination{}
	db, err2 := pagination.Paginate(c)
	if err1 != nil || err2 != nil {
		c.JSON(400, gin.H{"error": "invalid parameter."})
		return
	}
	fields := c.DefaultQuery("fields", "*")
	sel, db := setPreload(fields, db)

	var users []models.User
	err := db.Select(sel).Find(&users).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "error occured"})
		return
	}

	var index uint
	if len(users) < 1 {
		index = 0
	} else {
		index = users[len(users)-1].ID
	}
	pagination.SetHeaderLink(c, index)

	ver_range := semver.MustParseRange(">=1.0.0 <2.0.0")
	if ver_range(version) {
		// change the behavior depending on the version
		// 1.0.0 <= this < 2.0.0
	}

	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	version, err := middleware.VersionInit(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid parameter."})
		return
	}
	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	fields := c.DefaultQuery("fields", "*")
	sel, db := setPreload(fields, db)
	var user models.User
	err = db.Select(sel).First(&user, id).Error
	if err != nil {
		content := gin.H{"error": "user with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	ver_range := semver.MustParseRange(">=1.0.0 <2.0.0")
	if ver_range(version) {
		// change the behavior depending on the version
		// 1.0.0 <= this < 2.0.0
	}

	c.JSON(200, user)
	// curl -i http://localhost:8080/api/v1/users/1
}

func CreateUser(c *gin.Context) {
	version, err := middleware.VersionInit(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid parameter."})
		return
	}
	db := dbpkg.DBInstance(c)
	var user models.User
	c.Bind(&user)
	if db.Create(&user).Error != nil {
		content := gin.H{"error": "error occured"}
		c.JSON(500, content)
		return
	}
	ver_range := semver.MustParseRange(">=1.0.0 <2.0.0")
	if ver_range(version) {
		// change the behavior depending on the version
		// 1.0.0 <= this < 2.0.0
	}

	c.JSON(201, user)
}

func UpdateUser(c *gin.Context) {
	version, err := middleware.VersionInit(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid parameter."})
		return
	}
	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	var user models.User
	if db.First(&user, id).Error != nil {
		content := gin.H{"error": "user with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.Bind(&user)
	db.Save(&user)
	ver_range := semver.MustParseRange(">=1.0.0 <2.0.0")
	if ver_range(version) {
		// change the behavior depending on the version
		// 1.0.0 <= this < 2.0.0
	}

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	version, err := middleware.VersionInit(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid parameter."})
		return
	}
	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	var user models.User
	if db.First(&user, id).Error != nil {
		content := gin.H{"error": "user with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	db.Delete(&user)
	ver_range := semver.MustParseRange(">=1.0.0 <2.0.0")
	if ver_range(version) {
		// change the behavior depending on the version
		// 1.0.0 <= this < 2.0.0
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
