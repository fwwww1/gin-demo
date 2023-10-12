package main

import "fmt"

// const来定义枚举类型
// iota只能配合const()一起使用，iota只有在const进行累加效果
const (
	//可以在const()添加关键字 iota，每行的iota都会累加1，第一行的iota默认值是0
	BEIJING = iota
	SHANGHAI
	QINGDAO
)

func main() {
	//常量（只读）
	const length = 10

	fmt.Println("length=", length)
}
