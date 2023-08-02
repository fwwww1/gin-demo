package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//2.解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v", err)
		return
	}
	//3.渲染模板
	name := "zhangsan"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("render tempalte failed, err:%v", err)
	}
}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v\n", err)
		return
	}
}
