package redata

import (
	"time"
	"webserver/ider"
)

type Data struct {
	Status bool   `json:"status"` //状态
	Error  string `json:"error"`  //其他错误
	Date   string `json:"date"`
	Time   string `json:"time"`
	ID     int    `json:"id"`
	Hello  string `json:"hello"`
}

// Redata:返回一个结构体，这个结构体中包含了数据
func Redata() Data {

	var Data1 Data
	Data1.Status = true
	Data1.Error = "无"
	Data1.Date = time.Now().Format("2006:01:02") //日期
	Data1.Time = time.Now().Format("15:04:05")   //时间
	Data1.Hello = "你好！王召~"                       //文本
	Data1.ID = ider.GetID()
	//其他错误
	if false {
		//发生其他错误时候，将状态和错误写入结构体中
		Data1.Status = false
		Data1.Error = "其他错误！"
	}
	return Data1
}
