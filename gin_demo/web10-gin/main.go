package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//方法一：使用map
	data1 := map[string]interface{}{
		"name":    "zhangsan",
		"message": "hello",
		"age":     18,
	}
	r.GET("/json1", func(c *gin.Context) {
		c.JSON(http.StatusOK, data1)
	})
	//方法二：结构体,灵活使用tag对结构体字段做定制化操作
	type info struct {
		Name    string `json:"name"`
		Message string
		Age     int
	}
	data2 := info{
		"lisi",
		"hello",
		19,
	}
	r.GET("/json2", func(c *gin.Context) {
		c.JSON(http.StatusOK, data2)
	})
	r.Run()
}
