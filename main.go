package main

import (
	"go-user-api/db"
	"go-user-api/router"

	"github.com/gin-gonic/gin"
)

// main ...
func main() {
	r := SetupRouter()
	r.Run(":8080")
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	db.Connect(r)
	router.Initialize(r)
	return r
}
