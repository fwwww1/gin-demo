package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")
		//u := userInfo{
		//	username: username,
		//	password: password,
		//}
		var u userInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("userInformation:%v", u)
			c.JSON(http.StatusOK, gin.H{
				"Status": "ok",
			})
		}
	})

	r.POST("/form", func(c *gin.Context) {
		var u userInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("userInformation:%v", u)
			c.JSON(http.StatusOK, gin.H{
				"Status": "ok",
			})
		}
	})

	r.POST("/json", func(c *gin.Context) {
		var u userInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("userInformation:%v", u)
			c.JSON(http.StatusOK, gin.H{
				"Status": "ok",
			})
		}
	})
	r.Run()
}
