package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		//获取浏览器请求携带的query string参数
		name := c.Query("query")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	r.Run()
}
