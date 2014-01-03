package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type MyHandler struct {
}

func (p *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloName(w, r)
		return
	} else if r.URL.Path == "/login" {
		login(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //默认不解释参数
	fmt.Println(r.Form) //server端打印
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, " "))
	}
	fmt.Fprintf(w, "Hello Go Web，我是林秀全！") ////这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	//获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.tpl")
		err := t.Execute(w, nil)
		if err != nil {
			fmt.Println("表单执行错误：", err)
		}
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	/*
		//设置访问路由: 使用默认路由
		http.HandleFunc("/", sayHelloName)
		http.HandleFunc("/login", login)

		// Server监听端口
		err := http.ListenAndServe(":8080", nil)
	*/
	my := &MyHandler{}
	err := http.ListenAndServe(":8080", my)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
