package controllers

import (
	"net/http"
	"strings"

	dbpkg "github.com/shimastripe/gouserapi/db"
	"github.com/shimastripe/gouserapi/models"
	"github.com/shimastripe/gouserapi/version"

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
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	pagination := dbpkg.Pagination{}
	db, err := pagination.Paginate(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fields := c.DefaultQuery("fields", "*")
	sel, db := setPreload(fields, db)

	var users []models.User
	err = db.Select(sel).Find(&users).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "error occured"})
		return
	}

	// paging
	var index int
	if len(users) < 1 {
		index = 0
	} else {
		index = int(users[len(users)-1].ID)
	}
	pagination.SetHeaderLink(c, index)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(200, user)
	// curl -i http://localhost:8080/api/v1/users/1
}

func CreateUser(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(201, user)
}

func UpdateUser(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
