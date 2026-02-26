package main

import (
	"webserver/httpserver"
	"webserver/router"
)

// 定义处理函数

func main() {
	//注册路由
	router.Init()
	//监听端口
	httpserver.ListeningPort()

}
