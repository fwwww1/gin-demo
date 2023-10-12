package dao

import (
	"example/blog/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Manager interface {
	AddUser(user *model.User)
	Login(username string) model.User

	// AddPost 添加博客
	AddPost(post *model.Post)
	// GetAllPost 显示所有博客
	GetAllPost() []model.Post
	// GetPost 查询一个博客
	GetPost(pst int) model.Post
}
type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("faild open mysql,err : %v\n", err)
	}
	Mgr = &manager{db: db}
	//创建用户表
	db.AutoMigrate(&model.User{})
	//创建博客表
	db.AutoMigrate(&model.Post{})
}

func (mgr *manager) AddUser(user *model.User) {
	mgr.db.Create(user)
}
func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("username=?", username).First(&user)
	return user
}
func (mgr *manager) AddPost(post *model.Post) {
	mgr.db.Create(&post)
}
func (mgr *manager) GetAllPost() []model.Post {
	var posts = make([]model.Post, 10)
	mgr.db.Find(&posts)
	return posts
}
func (mgr *manager) GetPost(pst int) model.Post {
	var post model.Post
	mgr.db.First(&post, pst)
	return post
}
func main() {

}
