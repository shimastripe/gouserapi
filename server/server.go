package server

import (
	"github.com/shimastripe/gouserapi/db"
	"github.com/shimastripe/gouserapi/middleware"
	"github.com/shimastripe/gouserapi/router"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	DB := db.Connect(r)
	r.Use(middleware.SetDBtoContext(DB))
	router.Initialize(r)
	return r
}
