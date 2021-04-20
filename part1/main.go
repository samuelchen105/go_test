package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuhsuan105/go_test/common"
	"github.com/yuhsuan105/go_test/directory"
	"github.com/yuhsuan105/go_test/product"
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
	product.HandlerRegister(rt_api.Group("/pd"))

	rt.Run()
}
