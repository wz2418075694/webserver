package httpserver

import (
	"fmt"
	"net/http"
)

func ListeningPort() {

	//监听端口本机的8080端口
	fmt.Println("服务开启！")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println("连接失败！")
	}

}
