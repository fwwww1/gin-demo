package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//自定义一个函数
	//要么只有一个返回值，要么两个返回值，第二个必须是error类型
	kua := func(name string) (string, error) {
		return name + "真滴帅", nil
	}
	//定义模板
	t := template.New("f.tmpl") //创建一个名字为f的模板对象，名字一定要与模板的名字对应上
	//告诉模板引擎我现在自定义了一个函数，名字叫 kua
	t = t.Funcs(template.FuncMap{
		"kua": kua,
	})
	//解析模板
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v", err)
		return
	}
	//渲染模板
	name := "zhangsan"
	t.Execute(w, name)
}
func demo1(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v", err)
		return
	}
	name := "zhangsan"
	//渲染模板
	t.Execute(w, name)
}
func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo1)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v\n", err)
		return
	}
}
