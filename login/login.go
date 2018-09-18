package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func asyhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello astaxie")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		fmt.Printf("form的%v", r.Form)
		fmt.Println(len(r.Form.Get("username")), r.Form["username"], "eee")
		if len(r.Form.Get("username")) == 0 {
			fmt.Println("用户名不能为空")
			fmt.Fprintf(w, "用户名不能为空")
		} else {
			fmt.Fprintf(w, "登陆成功")
		}
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
func main() {
	http.HandleFunc("/", asyhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("LisntenAndServe:", err)
	}
}
