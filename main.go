package main

import (
	"github.com/shimastripe/gouserapi/db"
	"github.com/shimastripe/gouserapi/server"
)

// main ...
func main() {
	database := db.Connect()
	s := server.Setup(database)
	s.Run(":8080")
}
