package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//静态文件：html页面上用到的静态文件 .css js 图片

func main() {
	r := gin.Default()
	//加载静态页面
	r.Static("/xxx", "./statics")
	//gin框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//解析(加载)模板
	//r.LoadHTMLFiles("templates/index.tmpl","templates/users/index.tmpl")
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ //渲染模板
			"title": "baidu.com",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{ //渲染模板
			"title": "<a href='https://baidu.com'>百度</a>",
		})
	})

	//加载从网站上下载的模板
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})
	r.Run()
}
