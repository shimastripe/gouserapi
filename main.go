package main

import (
	"github.com/gin-gonic/gin"
)

// main ...
func main() {
	r := gin.Default()

	// 1. <a href="http://localhost:8080/" target="_blank" rel="noreferrer" style="cursor:help;display:inline !important;">http://localhost:8080/</a> へアクセスすると「Hello world」と表示する。
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello world")
	})

	//2. <a href="http://localhost:8080/hoge" target="_blank" rel="noreferrer" style="cursor:help;display:inline !important;">http://localhost:8080/hoge</a> へアクセスすると、「fuga」と表示する。
	r.GET("/hoge", func(c *gin.Context) {
		c.String(200, "fuga")
	})
	r.Run(":8080")
}
