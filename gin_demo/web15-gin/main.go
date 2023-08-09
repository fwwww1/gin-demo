package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", func(c *gin.Context) {
		//从请求中读取文件
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			//将获取到的文件保存到本地（服务端本地）
			//dst := fmt.Sprintf("./%s", file.Filename)
			dst := path.Join("./", file.Filename)
			c.SaveUploadedFile(file, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}

	})
	r.Run()
}
