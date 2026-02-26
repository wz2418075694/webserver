package main

import (
	"webserver/httpserver"
	"webserver/router"
)

// 定义处理函数

/*
1. 启动所有的其他模块，按顺序
2. 确保程序始终运行
*/
func main() {

	// 启动路由模块，注册路由
	router.Init()

	//启动httpserver模块，并监听端口
	if err := httpserver.Start("0.0.0.0:8080"); err != nil {
		panic(err)
	}

}
