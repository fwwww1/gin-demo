package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 1.定义模型
type User struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
}

func main() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:123456@(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//2.把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	/*//3.创建
	u1 := User{Name: "zhangsan", Age: 18, Active: true} //在代码层面创建一个User对象
	db.Create(&u1)
	u2 := User{Name: "jinzhu", Age: 20, Active: false} //在代码层面创建一个User对象
	db.Create(&u2)*/

	//4.查询
	var user User
	db.First(&user)

	/*//5.删除
	var u = User{}
	u.ID = 1
	db.Delete(&u)
	//db.Delete(User{},"age=?",18)*/
	/*var u1 []User
	db.Unscoped().Where("name=?", "wangwu").Find(&u1)
	fmt.Println(u1)*/

	//物理删除
	db.Debug().Unscoped().Where("name=?", "jinzhu").Delete(User{})
}
