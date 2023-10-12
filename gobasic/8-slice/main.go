package main

import "fmt"

func printArray(myArray []int) {
	//引用传递
	for _, value := range myArray {
		fmt.Println("value:", value)
	}
	myArray[0] = 100
}
func main() {
	//固定长度的数组
	var myArray1 [10]int

	myArray2 := [4]int{1, 2, 3, 4}

	myArray3 := []int{5, 6, 7, 8}
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}
	for index, value := range myArray2 {
		fmt.Println("index=", index, "value=", value)
	}

	printArray(myArray3)

	//声明slice是一个切片，并初始化，默认值是1 ，2 ，3，len为3
	//slice1:=[]int{1,2,3}
	//声明一个slice1切片，并没有分配空间
	//var slice1 []int
	//slice1 = make([]int, 10) //给slice分配10个内存空间，默认值是0
	//声明一个切片slice1并分配10个内存空间，初始化的值为0
	//var slice1 []int = make([]int,10)
	//声明一个切片slice1并分配10个内存空间，初始化的值为0，根据：=推导出slice1是一个切片
	//slice1 := make([]int, 10)

	//切片中追加元素
	var numbers = make([]int, 3, 5)
	fmt.Printf("len=%v,cap=%v,slice=%v", len(numbers), cap(numbers), numbers)

	//向numbers追加一个元素1,len=4 cap=5 , [0,0,0,1]
	numbers = append(numbers, 1)
	//向numbers追加一个元素5,len=5 cap=5 , [0,0,0,1,2]
	numbers = append(numbers, 2)
	//向一个容量已满的切片追加元素3，此时容量会变为原来的两倍
	numbers = append(numbers, 3)
	fmt.Printf("len=%v,cap=%v,slice=%v\n", len(numbers), cap(numbers), numbers)
	//切片中截取元素
	s := []int{1, 2, 3}
	//[0,2)
	s1 := s[0:2]
	fmt.Println(s1)
	//可以取到最后一个元素
	s2 := s[1:]
	fmt.Println(s2)
}
