package main

import "fmt"

func main() {
	//第一种声明方式
	//声明myMap1是一个map类型，key是string，value是string
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1是一个空map")
	}
	//分配空间
	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	fmt.Println(myMap1)

	//第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	fmt.Println(myMap2)

	//第三种声明方式
	myMap3 := map[string]string{
		"one": "python",
	}
	fmt.Println(myMap3)

	cityMap := make(map[string]string)
	cityMap["china"] = "beijing"
	cityMap["japan"] = "tokyo"
	cityMap["usa"] = "newyork"

	//遍历
	for key, value := range cityMap {
		fmt.Println("key:", key, "value:", value)
	}
	//删除
	delete(cityMap, "china")
	//修改
	cityMap["usa"] = "new"
	//遍历
	for key, value := range cityMap {
		fmt.Println("key:", key, "value:", value)
	}

}
