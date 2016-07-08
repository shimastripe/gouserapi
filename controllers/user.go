package controllers

import (
	"net/http"

	"github.com/shimastripe/gouserapi/db"
	"github.com/shimastripe/gouserapi/models"
	"github.com/shimastripe/gouserapi/pagination"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	pagination := pagination.Pagination{}
	db, err := pagination.Paginate(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid parameter."})
		return
	}
	fields := c.DefaultQuery("fields", "*")
	var users []models.User
	err = db.Select(fields).Preload("Profile").Preload("Profile.Nation").Preload("AccountName").Preload("Emails").Find(&users).Error
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

	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	fields := c.DefaultQuery("fields", "*")
	var user models.User
	err := db.Select(fields).Preload("Profile").Preload("Profile.Nation").Preload("AccountName").Preload("Emails").First(&user, id).Error
	if err != nil {
		content := gin.H{"error": "user with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.JSON(200, user)
	// curl -i http://localhost:8080/api/v1/users/1
}

func CreateUser(c *gin.Context) {
	db := db.DBInstance(c)
	var user models.User
	c.Bind(&user)
	if db.Create(&user).Error != nil {
		content := gin.H{"error": "error occured"}
		c.JSON(500, content)
		return
	}
	c.JSON(201, user)
}

func UpdateUser(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	var user models.User
	if db.First(&user, id).Error != nil {
		content := gin.H{"error": "user with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.Bind(&user)
	db.Save(&user)
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	var user models.User
	if db.First(&user, id).Error != nil {
		content := gin.H{"error": "user with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	db.Delete(&user)
	c.Writer.WriteHeader(http.StatusNoContent)
}
