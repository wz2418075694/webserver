package httpserver

import (
	"fmt"
	"net/http"
)

func Start(addr string) error {

	//监听端口本机的8080端口
	fmt.Println("服务开启！")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		return err
	}

	return nil
}
