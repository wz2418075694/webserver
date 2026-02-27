package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("D:\\GO\\webserver\\tmp\\zhansan.json")
	data2, _ := os.ReadFile("D:\\GO\\webserver\\tmp\\wangzhao.json")
	fmt.Println(string(data2))
	fmt.Println(string(data))

}
