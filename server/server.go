package server

import (
	"go-user-api/db"
	"go-user-api/router"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	db.Connect(r)
	router.Initialize(r)
	return r
}
