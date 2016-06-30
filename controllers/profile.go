package controllers

import (
	"net/http"
	"strings"

	"github.com/shimastripe/gouserapi/db"
	"github.com/shimastripe/gouserapi/models"
	"github.com/shimastripe/gouserapi/query"

	"github.com/gin-gonic/gin"
)

func GetProfiles(c *gin.Context) {
	db := db.DBInstance(c)
	var profiles []models.Profile
	db.Find(&profiles)
	c.JSON(200, profiles)
}

func GetProfile(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	fields := c.DefaultQuery("fields", "")
	var profile models.Profile
	var err error
	if fields != "" {
		err = db.Select(fields).First(&profile, id).Error
	} else {
		err = db.First(&profile, id).Error
	}

	if err != nil {
		content := gin.H{"error": "profile with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.JSON(200, query.FilterField(strings.Split(fields, ","), &profile))
	// curl -i http://localhost:8080/api/v1/profiles/1
}

func CreateProfile(c *gin.Context) {
	db := db.DBInstance(c)
	var profile models.Profile
	c.Bind(&profile)
	err := db.Create(&profile).Error
	if err != nil {
		content := gin.H{"error": err}
		c.JSON(500, content)
		return
	}
	c.JSON(201, profile)
}

func UpdateProfile(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	var profile models.Profile
	if db.First(&profile, id).Error != nil {
		content := gin.H{"error": "profile with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.Bind(&profile)
	db.Save(&profile)
	c.JSON(200, profile)
}

func DeleteProfile(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	var profile models.Profile
	if db.First(&profile, id).Error != nil {
		content := gin.H{"error": "profile with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	db.Delete(&profile)
	c.Writer.WriteHeader(http.StatusNoContent)
}
