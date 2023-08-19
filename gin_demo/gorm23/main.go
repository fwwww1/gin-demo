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

	//5.更新
	user.Name = "张三"
	user.Age = 99
	db.Save(&user) //默认会修改所有字段

	db.Debug().Model(&user).Update("name", "lisi")

	m1 := map[string]interface{}{
		"name":   "wangwu",
		"age":    28,
		"active": true,
	}
	db.Debug().Model(&user).Updates(m1)                //m1列出来的字段都会更新
	db.Debug().Model(&user).Select("age").Updates(m1)  //只更新m1中的age字段
	db.Debug().Model(&user).Omit("active").Updates(m1) //排除m1中的 active 更新其它字段

	//让user表中所有用户的年龄+2
	db.Model(&User{}).Update("age", gorm.Expr("age+?", 2))
}
