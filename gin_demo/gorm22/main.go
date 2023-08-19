package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 1.定义模型
type User struct {
	gorm.Model
	Name string
	Age  int64
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

	//3.创建
	//u1 := User{Name: "zhangsan", Age: 18} //在代码层面创建一个User对象
	//db.Create(&u1)
	//u2 := User{Name: "jinzhu", Age: 20} //在代码层面创建一个User对象
	//db.Create(&u2)

	//4.查询
	//var user User
	//user:=new(User)
	//db.First(&user)
	//fmt.Printf("user:%v", user)

	//var users []User
	//db.Find(&users)
	//fmt.Printf("users:%v", users)

	//FirstOrInit
	var user User
	//db.Attrs(User{Age: 99}).FirstOrInit(&user, User{Name: "lisi"}) //Attrs 查询不到时初始化参数
	db.Assign(User{Age: 99}).FirstOrInit(&user, User{Name: "zhangsan"}) //Assign 不管是否查询到记录，都将参数赋值给struct
	fmt.Printf("user:%v", user)

	//Select 选择字段，指定你想从数据库中检索出的字段
	var users []User
	db.Select("name,age").Find(&users)
	fmt.Printf("users:%v", users)
}
