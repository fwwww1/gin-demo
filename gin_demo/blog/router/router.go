package router

import (
	"example/blog/controller"
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	//解析模板
	r.LoadHTMLGlob("templates/*")
	//静态文件
	r.Static("../assets", "./assets")

	r.GET("/", controller.Index)              //首页
	r.GET("/register", controller.GoRegister) //跳转到注册页面
	r.POST("/register", controller.Register)  //进行注册
	r.GET("/login", controller.GoLogin)       //进入登录页面
	r.POST("/login", controller.Login)        //登录

	//操作博客
	//博客列表
	r.POST("/post_index", controller.GetPostIndex)
	//添加博客
	r.POST("/post", controller.AddPost)
	//跳转到添加博客页面
	r.GET("/post", controller.GoAddPost)
	r.Run()
}
func main() {

}
