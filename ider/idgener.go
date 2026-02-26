package ider

import (
	"math/rand"
	"time"
)

var id int

func init() {
	rand.Seed(time.Now().UnixNano())
	id = rand.Intn(100)
	//启动定时器
	go func() {
		//60秒触发一次
		ticker := time.NewTicker(60 * time.Second)
		//循环监听定时器
		for range ticker.C {
			id = rand.Intn(100) //更新全局变量
		}
	}()
}

// idgener就是一个通用的ID生成器， 对外暴露GetID用来获取当前ID值
func GetID() int {
	//fmt.Println(id)
	return id
}
