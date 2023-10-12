package main

import "fmt"

// 定义一个Book结构体

type Book struct {
	Title string
	auth  string
}

func changeBook1(book Book) {
	//传递副本
	book.auth = "lisi"
}
func changeBook2(book *Book) {
	book.auth = "666"
}

func main() {
	//声明一个Book结构体类型的实例
	var book1 Book
	book1.Title = "Golang"
	book1.auth = "zhangsan"
	fmt.Println(book1)
	changeBook1(book1)
	fmt.Println(book1)
	changeBook2(&book1)
	fmt.Println(book1)

}
