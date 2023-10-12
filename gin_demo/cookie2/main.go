package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取客户端cookie并校验
		cookie, err := c.Cookie("abc")
		if err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		//返回错误信息
		c.JSON(http.StatusOK, gin.H{"error": err})
		//若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}
func main() {
	r := gin.Default()
	//login
	r.GET("/login", func(c *gin.Context) {
		//设置cookie
		c.SetCookie("abc", "123", 3600, "/", "localhost", false, true)
		//返回信息
		c.JSON(http.StatusOK, "Login success")
	})
	//home
	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "home"})
	})
	r.Run()
}
