package controller

import (
	"example/blog/dao"
	"example/blog/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//显示用户

func ListUser(c *gin.Context) {
	c.HTML(http.StatusOK, "userList.html", nil)
}

//首页

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

//注册页面

func GoRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

//注册用户

func Register(c *gin.Context) {
	//获取用户数据
	username := c.PostForm("username")
	password := c.PostForm("password")
	//存入结构体
	user := model.User{
		Username: username,
		Password: password,
	}
	//创建记录
	dao.Mgr.AddUser(&user)
	c.Redirect(http.StatusMovedPermanently, "/")
}

//登录页面

func GoLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

//登录

func Login(c *gin.Context) {
	//获取用户登录时提交的用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")
	//在数据库中查询用户名
	user := dao.Mgr.Login(username)
	//判断用户是否存在
	if user.Username == username {
		//若存在，判断密码是否正确
		if user.Password == password {
			fmt.Println("登录成功")
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			c.HTML(http.StatusOK, "login.html", "密码错误")
		}
	} else {
		c.HTML(http.StatusOK, "login.html", "用户名不存在")
	}
}

// 博客列表

func GetPostIndex(c *gin.Context) {
	//查询所有博客
	posts := dao.Mgr.GetAllPost()
	//内容渲染到页面
	c.HTML(http.StatusOK, "postIndex.html", posts)
}

//添加博客

func AddPost(c *gin.Context) {
	//获取表单内容
	title := c.PostForm("title")
	content := c.PostForm("content")
	tag := c.PostForm("tag")

	var post = model.Post{
		Title:   title,
		Content: content,
		Tag:     tag,
	}
	dao.Mgr.AddPost(&post)
	c.Redirect(http.StatusFound, "/post_index")
}

//跳转到添加博客页面

func GoAddPost(c *gin.Context) {
	c.HTML(http.StatusOK, "post.html", nil)
}
func main() {

}
