package main

import "fmt"

func main() {
	var (
		name    string
		age     int
		married bool
		num     []int
	)
	//fmt.Scan(&name, &age, &married, &num)
	//fmt.Println(name, age, married)
	fmt.Scanf("1:%s,2:%d,3:%t,4:%v", &name, &age, &married, &num)
	fmt.Println(name, age, married, num)
}
