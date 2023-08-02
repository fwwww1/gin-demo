package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	file, _ := ioutil.ReadFile("hello.txt")
	_, _ = fmt.Fprintf(w, string(file))
}
func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Print("http serve failed, err:%v\n", err)
		return
	} else {
		fmt.Print("serve success")
	}
}
