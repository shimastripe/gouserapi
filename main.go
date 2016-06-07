package main

import (
	"go-user-api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

// main ...
func main() {
	var err error
	db, err = gorm.Open("sqlite3", "db/users.db")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	db.AutoMigrate(&models.User{})

	r := gin.Default()
	api := r.Group("api")
	{
		api.GET("/users", GetUsers)
		api.GET("/users/:id", GetUser)
		api.POST("/users", PostUser)
		api.PUT("/users/:id", UpdateUser)
		api.DELETE("/users/:id", DeleteUser)
	}
	r.Run(":8080")
}

func GetUsers(c *gin.Context) {
	var users []models.User
	db.Find(&users)
	c.JSON(200, users)
	// curl -i http://localhost:8080/api/v1/users
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	if db.First(&user, id).Error != nil {
		content := gin.H{"error": "user with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.JSON(200, user)
	// curl -i http://localhost:8080/api/v1/users/1
}
func PostUser(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	db.Create(&user)
	c.JSON(201, user)
}
func UpdateUser(c *gin.Context) {
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
