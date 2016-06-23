package main

import "github.com/shimastripe/gouserapi/server"

// main ...
func main() {
	r := server.SetupRouter()
	r.Run(":8080")
}
