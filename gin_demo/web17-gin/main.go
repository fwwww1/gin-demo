package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//Any：请求方法的集合
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"method": "POST",
			})
		}
	})
	//NoRouter  访问没有定义路由的地址
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "zhangsan"})
	})
	//路由组 多用于区分不同业务线或API版本
	//把公用的前缀提取出来，创建一个路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/index",
			})
		})
		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/xx",
			})
		})
	}
	r.Run()
}
