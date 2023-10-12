package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	return c
}

// 多个返回值，匿名
func fool2(a string, b int) (int, int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	return 600, 700
}

// 多个返回值，有形参名称
func fool3(a string, b int) (r1 int, r2 int) {
	fmt.Println("==fool3==")
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	//r1 r2 foo3的形参，初始化默认值是0
	//r1 r2 作用域空间 是foo3 整个函数体{}空间
	fmt.Println("r1=", r1)
	fmt.Println("r2=", r2)
	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}
func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("==fool4==")
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}
func main() {
	c := foo1("abc", 1)
	fmt.Println("c=", c)

	ret1, ret2 := fool2("def", 2)
	fmt.Println("ret1=", ret1, "ret2=", ret2)

	r1, r2 := fool3("ghi", 3)
	fmt.Println("r1=", r1, "r2=", r2)

	r1, r2 = foo4("jkl", 4)
	fmt.Println("r1=", r1, "r2=", r2)
}
