package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuhsuan105/go_test/common"
	"github.com/yuhsuan105/go_test/directory"
)

func main() {
	common.InitDatabase()

	rt := gin.Default()
	rt.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	rt_api := rt.Group("/api")
	directory.HandlerRegister(rt_api.Group("/dir"))

	rt.Run()
}