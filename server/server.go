package server

import (
	"github.com/shimastripe/gouserapi/db"
	"github.com/shimastripe/gouserapi/router"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	db.Connect(r)
	router.Initialize(r)
	return r
}
