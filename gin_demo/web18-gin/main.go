package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//中间件 middleware

// 定义一个中间件m1统计请求处理函数耗时
func m1(c *gin.Context) {
	fmt.Println("m1...in...")
	//go funcXX(c.Copy()) //在funcXX中只能使用c的拷贝
	//计时
	start := time.Now()
	c.Next() //调用后续的处理函数
	//c.Abort() //阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v", cost)
	fmt.Println("m1...out...")
}
func m2(c *gin.Context) {
	fmt.Println("m2...in...")
	c.Set("name", "zhangsan")
	//c.Next() //调用后续的处理函数
	//c.Abort() //阻止调用后续的处理函数
	fmt.Println("m2...out...")
}
func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//或者一些其它准备工作
	return func(c *gin.Context) {
		if doCheck {
			//存放具体逻辑
		} else {
			c.Next()
		}
	}
}
func main() {
	r := gin.Default()                  //默认使用了Logger()和Recovery()中间件
	r.Use(m1, m2, authMiddleware(true)) //全局注册中间件函数
	r.GET("/index1", func(c *gin.Context) {
		name, ok := c.Get("name")
		if !ok {
			name = "匿名用户"
		}
		c.JSON(http.StatusOK, gin.H{
			"message": name,
		})
	})
	/*//路由组注册中间件方法一
	xxGroup := r.Group("/xx", authMiddleware(true))
	{
		xxGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "xxGroup"})
		})
	}
	//路由组注册中间件方法二
	xx2Group := r.Group("/xx")
	xx2Group.Use(authMiddleware(true))
	{
		xx2Group.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "xx2Group"})
		})
	}*/
	r.Run()
}
