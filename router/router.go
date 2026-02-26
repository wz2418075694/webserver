package router

import (
	"net/http"
	"webserver/httpserver"
	"webserver/studenter"
)

func Init() {
	//注册路由函数
	http.HandleFunc("/hello", httpserver.Hello)
	http.HandleFunc("/time", httpserver.Time)
	http.HandleFunc("/student", studenter.StudentHandler)
}
