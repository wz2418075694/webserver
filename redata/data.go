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

//var lastUpdateTime time.Time

// 写一个初始化函数，在程序跑之前开始执行，且执行一次
//
//	func init() {
//		//生成随机数
//		//Data1.ID = ider.GetID()
//		//lastUpdateTime = time.Now()
//
//		//初始化默认状态
//	}
func Redata() Data {
	var Data1 Data
	Data1.Status = true
	Data1.Error = "无"
	Data1.Date = time.Now().Format("2006:01:02")
	Data1.Time = time.Now().Format("15:04:05")
	Data1.Hello = "你好！王召~"
	Data1.ID = ider.GetID()
	//其他错误
	if false {
		Data1.Status = false
		Data1.Error = "其他错误！"
	}
	return Data1
}

//NowTime := time.Now()
//if (NowTime.Sub(lastUpdateTime)).Seconds() > 60 {
//	//更新随机数
//	Data1.ID = ider.GetID()
//	lastUpdateTime = NowTime
//	fmt.Println("id已经更新！", Data1.ID) //看id变化
//}
