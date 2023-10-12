package main

import "fmt"

// interface{}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called ...")
	fmt.Println(arg)

	//interface{}如何区分此时引用的底层数据类型是什么
	//给 interface{}提供类型”断言机制“
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value=", value)
		fmt.Printf("value type is %T/n", value)
	}
}

type book struct {
	auth string
}

func main() {
	b := book{"Golang"}
	myFunc(b)
	myFunc("abc")
	myFunc(1)
}
