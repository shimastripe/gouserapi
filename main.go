package main

import "go-user-api/server"

// main ...
func main() {
	r := server.SetupRouter()
	r.Run(":8080")
}
