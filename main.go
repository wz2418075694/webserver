package main

import (
	"github.com/HouGuoFa/golog"
	"webserver/httpserver"
	"webserver/mysql"
	"webserver/router"
)

/*
1. 启动所有的其他模块，按顺序
2. 确保程序始终运行
*/
func main() {

	// 启动路由模块，注册路由
	router.Init()

	if err := mysql.InitDB(); err != nil {
		golog.Error(err)
	}

	//启动httpserver模块，并监听端口
	if err := httpserver.Start("0.0.0.0:8080"); err != nil {
		panic(err) //程序崩溃，打印panic函数中的东西
	}

}
