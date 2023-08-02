package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v", err)
		return
	}
	//渲染模板
	u1 := User{
		Name:   "zhangsan",
		Gender: "男",
		Age:    18,
	}
	m1 := map[string]interface{}{
		"name":   "lisi",
		"gender": "男",
		"age":    18,
	}
	hobbyList := []string{
		"睡觉",
		"吃饭",
		"打游戏",
	}
	t.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})
}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v\n", err)
		return
	}
}
