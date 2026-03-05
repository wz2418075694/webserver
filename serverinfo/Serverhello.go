package serverinfo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webserver/redata"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	//序列化json数组，将对象转化为字符串
	jsonbytes, err := json.Marshal(redata.Redata())
	if err != nil { //序列化失败
		http.Error(w, "结构体序列化失败！"+err.Error(), 500)
		return
	}
	fmt.Println("json结构体序列化成功！")
	//设置响应头
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//发送json字节数据
	_, err = w.Write(jsonbytes)
	if err != nil {
		fmt.Println("发送失败！")
		http.Error(w, "发送失败"+err.Error(), 500)
		return
	}
}
