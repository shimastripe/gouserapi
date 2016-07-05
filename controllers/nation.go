package controllers

import (
	"net/http"

	"github.com/shimastripe/gouserapi/db"
	"github.com/shimastripe/gouserapi/models"

	"github.com/gin-gonic/gin"
)

func GetNations(c *gin.Context) {
	db := db.DBInstance(c)
	fields := c.DefaultQuery("fields", "*")
	var nations []models.Profile
	db.Select(fields).Find(&nations)
	c.JSON(200, nations)
}

func GetNation(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	fields := c.DefaultQuery("fields", "*")
	var nation models.Nation
	err := db.Select(fields).First(&nation, id).Error
	if err != nil {
		content := gin.H{"error": "nation with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.JSON(200, &nation)
	// curl -i http://localhost:8080/api/v1/nations/1
}

func CreateNation(c *gin.Context) {
	db := db.DBInstance(c)
	var nation models.Nation
	c.Bind(&nation)
	if db.Create(&nation).Error != nil {
		content := gin.H{"error": "error occured"}
		c.JSON(500, content)
		return
	}
	c.JSON(201, nation)
}

func UpdateNation(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	var nation models.Nation
	if db.First(&nation, id).Error != nil {
		content := gin.H{"error": "nation with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.Bind(&nation)
	db.Save(&nation)
	c.JSON(200, nation)
}

func DeleteNation(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	var nation models.Nation
	if db.First(&nation, id).Error != nil {
		content := gin.H{"error": "nation with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	db.Delete(&nation)
	c.Writer.WriteHeader(http.StatusNoContent)
}
