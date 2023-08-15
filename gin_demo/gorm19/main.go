package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// gorm demo1

// UserInfo --> 数据表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:123456@(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	//创建表  自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{})

	//创建数据行
	//u1 := UserInfo{1, "zhangsan", "男", "game"}
	//db.Create(&u1)

	//查询
	var u UserInfo
	db.First(&u) //查询表中第一条数据
	fmt.Printf("first record:%v", u)

	//更新
	db.Model(&u).Update("hobby", "lol")

	//删除
	db.Delete(&u)

	defer db.Close()
}
