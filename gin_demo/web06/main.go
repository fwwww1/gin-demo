package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.New("f").ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v", err)
		return
	}
	//渲染模板
	name := "zhangsan"
	t.Execute(w, name)
}
func main() {
	http.HandleFunc("/", f1)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v\n", err)
		return
	}
}
