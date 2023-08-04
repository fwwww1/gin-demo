package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v", err)
		return
	}
	msg := "这是index"
	//渲染模板
	t.Execute(w, msg)
}
func home(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v", err)
		return
	}
	msg := "这是home"
	//渲染模板
	t.Execute(w, msg)
}

func index2(w http.ResponseWriter, r *http.Request) {
	//定义模板（模板继承的方式）
	//解析模板
	t, err := template.ParseFiles("./template/base.tmpl", "./template/index2.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v", err)
		return
	}
	//渲染模板
	msg := "这是index2模板"
	t.ExecuteTemplate(w, "index2.tmpl", msg)
}
func home2(w http.ResponseWriter, r *http.Request) {
	//定义模板（模板继承的方式）
	//解析模板
	t, err := template.ParseFiles("./template/base.tmpl", "./template/home2.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v", err)
		return
	}
	//渲染模板
	msg := "这是home2模板"
	t.ExecuteTemplate(w, "home2.tmpl", msg)
}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v\n", err)
		return
	}
}
