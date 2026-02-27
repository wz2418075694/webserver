package serverinfo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webserver/redata"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	//序列化json数组
	jsonDate, err := json.Marshal(redata.Redata())
	if err != nil { //序列化失败
		http.Error(w, "结构体序列化失败！", 500)
		return
	}
	fmt.Println("json序列化成功！")
	//设置响应头
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//发送json字节数据
	_, err = w.Write(jsonDate)
	if err != nil {
		fmt.Println("发送失败！")
	}
}
