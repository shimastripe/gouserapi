package controllers

import (
	"net/http"
	"strings"

	"github.com/shimastripe/gouserapi/db"
	"github.com/shimastripe/gouserapi/models"
	"github.com/shimastripe/gouserapi/query"

	"github.com/gin-gonic/gin"
)

func GetAccountNames(c *gin.Context) {
	db := db.DBInstance(c)
	var account_names []models.AccountName
	db.Find(&account_names)
	c.JSON(200, account_names)
}

func GetAccountName(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	fields := c.DefaultQuery("fields", "")
	var account_name models.AccountName
	var err error
	if fields != "" {
		err = db.Select(fields).First(&account_name, id).Error
	} else {
		err = db.First(&account_name, id).Error
	}

	if err != nil {
		content := gin.H{"error": "account_name with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.JSON(200, query.FilterField(strings.Split(fields, ","), &account_name))
	// curl -i http://localhost:8080/api/v1/account_names/1
}

func CreateAccountName(c *gin.Context) {
	db := db.DBInstance(c)
	var account_name models.AccountName
	c.Bind(&account_name)
	err := db.Create(&account_name).Error
	if err != nil {
		content := gin.H{"error": err}
		c.JSON(500, content)
		return
	}
	c.JSON(201, account_name)
}

func UpdateAccountName(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	var account_name models.AccountName
	if db.First(&account_name, id).Error != nil {
		content := gin.H{"error": "account_name with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.Bind(&account_name)
	db.Save(&account_name)
	c.JSON(200, account_name)
}

func DeleteAccountName(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	var account_name models.AccountName
	if db.First(&account_name, id).Error != nil {
		content := gin.H{"error": "account_name with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	db.Delete(&account_name)
	c.Writer.WriteHeader(http.StatusNoContent)
}
