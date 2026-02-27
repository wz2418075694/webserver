package studenter

import (
	"fmt"
	"net/http"
	"os"
)

// 定义存储目录（固定路径，所有学生文件都存在这里）
const storageDir = "D:\\GO\\webserver\\tmp"

// init函数：程序启动时自动执行，做初始化工作
func init() {
	// MkdirAll：创建目录，0755是目录权限（可读可写）
	// 即使目录已存在，也不会报错，很安全
	if err := os.MkdirAll(storageDir, 0755); err != nil {
		// panic：如果目录创建失败，直接终止程序（没地方存文件，程序没法运行）
		panic(fmt.Sprintf("创建存储目录失败: %v", err))
	}
}

// student模块入库函数
func StudentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		doRead(w, r)
	case http.MethodPost:
		doWrite(w, r)
	}
}
