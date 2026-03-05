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
		//创建一个定时器，在后台监听60秒触发一次，往自己的c通道里面发送一个当前的时间
		ticker := time.NewTicker(60 * time.Second)
		//循环监听定时器
		for range ticker.C { //如果自己的通道里面有时间了，就执行一下随机数函数
			id = rand.Intn(100) //更新全局变量
		}
	}()

	//go checker()
}

// idgener就是一个通用的ID生成器， 对外暴露GetID用来获取当前ID值
func GetID() int {
	//fmt.Println(id)
	return id
}

//func checker() {
//	tick := time.Tick(time.Second)
//	tick2 := time.NewTicker(5 * time.Second)
//
//	for {
//		select {
//		case <-tick:
//			fmt.Println(time.Now(), GetID())
//		case <-tick2.C:
//			fmt.Println(time.Now(), "i'm ok ！")
//		}
//	}
//}
