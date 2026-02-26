package router

import (
	"net/http"
	"webserver/serverinfo"
	"webserver/studenter"
)

func Init() {
	//注册路由函数
	http.HandleFunc("/hello", serverinfo.Hello)
	http.HandleFunc("/time", serverinfo.Time)
	http.HandleFunc("/student", studenter.StudentHandler)
}
