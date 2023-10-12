package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			//退出当前goroutine
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}

}
