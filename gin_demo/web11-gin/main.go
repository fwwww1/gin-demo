package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//querystring

	//GET请求，URL ？后面的是query string参数
	//key=value格式，多个key-value用&连接
	r.GET("/web", func(c *gin.Context) {
		//获取浏览器请求携带的query string参数
		//name := c.Query("query") //通过Query获取请求中携带的query string参数
		//name := c.DefaultQuery("query", "somebody")//取不到就用默认值
		name, ok := c.GetQuery("query") //取到返回（值，true），取不到返回（"",false）
		if !ok {
			name = "somebody"
		}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	r.Run()
}
