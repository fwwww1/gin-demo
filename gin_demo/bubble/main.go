package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {
	//创建数据库
	//sql: create database bubble;
	//连接数据库
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//模型绑定
	db.AutoMigrate(&Todo{})
	//关闭数据库
	defer db.Close()
	//路由引擎
	r := gin.Default()
	//静态文件 (告诉模板文件引用的静态文件去哪找)
	r.Static("/static", "static")
	//解析index.html
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		//渲染index.html
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//v1
	v1Group := r.Group("/v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo Todo
			//前端页面填写待办事项，点击提交，会发送请求到这里
			//1.从请求中把数据拿过来
			c.BindJSON(&todo)
			//2.将数据存入数据库
			err2 := db.Create(&todo).Error
			//3.返回响应
			if err2 != nil {
				c.JSON(http.StatusOK, gin.H{"err:": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//查看（显示）所有待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			//查询todos这个表中的所有数据
			var todoList []Todo
			if err2 := db.Find(&todoList).Error; err2 != nil {
				c.JSON(http.StatusOK, gin.H{"err:": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		//查看一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {
			var todo Todo
			//获取查询待办事项的ID
			//c.Params.Get("id")
			id := c.Param("id")
			fmt.Println(id)
			//从数据库中查询id=？的数据
			if err2 := db.Where("id=?", id).First(&todo).Error; err2 != nil {
				c.JSON(http.StatusOK, gin.H{"err:": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			var todo Todo
			//获取要修改的待办事项ID
			//c.Params.Get("id")
			id := c.Param("id")
			fmt.Println(id)
			//从数据库中查询id=？的数据
			if err := db.Debug().Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"err:": err.Error()})
				return
			}
			//请求中修改后的信息保存到todo
			c.BindJSON(&todo)
			//将修改后的信息保存到数据库
			if err := db.Debug().Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"err:": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			//获取要删除待办事项的ID
			id := c.Param("id")
			//从数据库中删除id=？的数据
			if err := db.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"err:": err.Error()})
				return
			}
		})
	}

	r.Run()
}
