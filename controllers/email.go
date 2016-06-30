package controllers

import (
	"net/http"
	"strings"

	"github.com/shimastripe/gouserapi/db"
	"github.com/shimastripe/gouserapi/models"
	"github.com/shimastripe/gouserapi/query"

	"github.com/gin-gonic/gin"
)

func GetEmails(c *gin.Context) {
	db := db.DBInstance(c)
	var emails []models.Email
	db.Find(&emails)
	c.JSON(200, emails)
}

func GetEmail(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	fields := c.DefaultQuery("fields", "")
	var email models.Email
	var err error
	if fields != "" {
		err = db.Select(fields).First(&email, id).Error
	} else {
		err = db.First(&email, id).Error
	}

	if err != nil {
		content := gin.H{"error": "email with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.JSON(200, query.FilterField(strings.Split(fields, ","), &email))
	// curl -i http://localhost:8080/api/v1/emails/1
}

func CreateEmail(c *gin.Context) {
	db := db.DBInstance(c)
	var email models.Email
	c.Bind(&email)
	if db.Create(&email).Error != nil {
		content := gin.H{"error": "error occured"}
		c.JSON(500, content)
		return
	}
	c.JSON(201, email)
}

func UpdateEmail(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	var email models.Email
	if db.First(&email, id).Error != nil {
		content := gin.H{"error": "email with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.Bind(&email)
	db.Save(&email)
	c.JSON(200, email)
}

func DeleteEmail(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	var email models.Email
	if db.First(&email, id).Error != nil {
		content := gin.H{"error": "email with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	db.Delete(&email)
	c.Writer.WriteHeader(http.StatusNoContent)
}
