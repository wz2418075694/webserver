package httpserver

import (
	"fmt"
	"net/http"
)

func Start(addr string) error {
	//将这里变成手动传入地址，返回一个错误类型

	//监听端口本机的8080端口
	fmt.Println("服务开启！")
	err := http.ListenAndServe(addr, nil)
	//nil	用Go内置的 “默认大堂经理”（DefaultServeMux），按 HandleFunc 注册的路径分配请求
	//自定义handler完全接管所有请求，请求怎么处理、分配，全由你写的这个handler决定
	if err != nil {
		return err
	}
	return nil
}
