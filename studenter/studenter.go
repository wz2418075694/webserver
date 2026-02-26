package studenter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// student模块入库函数
func StudentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		doRead(w, r)
	case http.MethodPost:
		doWrite(w, r)
	}
}

// 处理读取数据
// http://127.1:8080/student?name=wangzhao
func doRead(w http.ResponseWriter, r *http.Request) {

	name := strings.TrimSpace(r.FormValue("name"))
	fmt.Println(r.URL.String(), name)
	if len(name) <= 0 { // 取全部信息
		return
	}

	// 取单个信息
	doSingleStudent(w, r)

}

// 处理创建传入数据
func doWrite(w http.ResponseWriter, r *http.Request) {

	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println("收到新信息！")

	filename := genFileFullName("")
	if err = dumpFile(filename, data); err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
}

func genFileFullName(file string) string {
	return "D:\\GO\\webserver\\tmp\\wangzhao"
}

// 写文件
func dumpFile(file string, content []byte) error {
	return os.WriteFile(file, content, 0x777)
}

// 处理单个学生信息查询
func doSingleStudent(w http.ResponseWriter, r *http.Request) {

	//student, err := readStudentFromJSON("D:\\GO\\webserver\\student.json")
	student, err := readStudentFromJSON(genFileFullName(""))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err = sendJson(w, student); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("send succ ！")

}

func sendJson(w http.ResponseWriter, Sdata any) error {

	data, err := json.Marshal(Sdata)
	if err != nil { //序列化失败
		http.Error(w, "结构体序列化失败！", 500)
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = w.Write(data)
	return err
}

// 读取JSON文件并解析为单个学生对象
func readStudentFromJSON(jsonfile string) (*Student, error) {

	// 读取文件内容
	data, err := os.ReadFile(jsonfile)
	if err != nil {
		return nil, fmt.Errorf("读取文件内容失败: %v", err)
	}
	// 解析JSON数据到Student结构体,将文本对象转化为内存对象
	var student *Student
	err = json.Unmarshal(data, student)
	if err != nil {
		return nil, fmt.Errorf("解析JSON数据失败: %v", err)
	}
	return student, nil
}
