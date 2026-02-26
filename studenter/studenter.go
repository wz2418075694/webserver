package studenter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Student struct {
	Name   string `json:"name"`   // 姓名
	Gender string `json:"gender"` // 性别
	Age    int    `json:"age"`    // 年龄
	Phone  string `json:"phone"`  // 电话
	City   string `json:"city"`   // 城市
}

// 读取JSON文件并解析为单个学生对象
func readStudentFromJSON(jsonfile string) (Student, error) {
	// 读取文件内容
	data, err := os.ReadFile(jsonfile)
	fmt.Println(jsonfile)
	//content, err := os.ReadFile(jsonfile)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(content))
	if err != nil {
		return Student{}, fmt.Errorf("读取文件内容失败: %v", err)
	}
	// 解析JSON数据到Student结构体,将文本对象转化为内存对象
	var student Student
	err = json.Unmarshal(data, &student)
	if err != nil {
		return Student{}, fmt.Errorf("解析JSON数据失败: %v", err)
	}
	return student, nil
}
func StudentHandler(w http.ResponseWriter, r *http.Request) {
	//调用函数
	student, err := readStudentFromJSON("student.json")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	studentdata, err := json.Marshal(student)
	if err != nil { //序列化失败
		http.Error(w, "结构体序列化失败！", 500)
		return
	}
	fmt.Println("json序列化成功！")
	//设置响应头
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//发送json字节数据
	_, err = w.Write(studentdata)
	if err != nil {
		fmt.Println("发送失败！")
	}

}
