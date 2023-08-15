package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 1.定义模型
type User struct {
	ID   int64
	Name *string `gorm:"default:'zhangsan'"`
	//Name sql.NullString `gorm:"default:'zhangsan'"`
	Age int64
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
	u := User{Name: new(string), Age: 18} //在代码层面创建一个User对象
	//u := User{Name: sql.NullString{"",true}, Age: 18} //在代码层面创建一个User对象
	fmt.Println(db.NewRecord(&u)) //判断主键是否为空
	db.Create(&u)
	fmt.Println(db.NewRecord(&u)) //判断主键是否为空

}
