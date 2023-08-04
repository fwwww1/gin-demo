package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.New("index.tmpl").Delims("{[", "]}").ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v", err)
		return
	}
	//渲染模板
	name := "zhangsan"
	t.Execute(w, name)
}
func xss(w http.ResponseWriter, r *http.Request) {
	//定义模板

	//解析模板
	//解析模板之前定义一个"safe"函数
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v", err)
		return
	}
	str1 := "<script>alert(123);</script>"
	str2 := "<a href='https://baidu.com'>百度</a>"
	//渲染模板
	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v\n", err)
		return
	}
}
