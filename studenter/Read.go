package studenter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var name string

// 处理读取数据
func doRead(w http.ResponseWriter, r *http.Request) {

	//strings.TrimSpace函数的作用是删除字符串开头和结尾所用的空格
	//r.FormValue的作用是在http请求中提取name的值，如果没有返回空字符串
	name = strings.TrimSpace(r.FormValue("name"))
	fmt.Println(r.URL.String(), name)
	if len(name) <= 0 { // 取全部信息
		doAllstudents(w, r)
		//return
	} else { // 取单个信息
		doSingleStudent(w, r)
	}
	return

}
func getFilePath(file string) string {
	filePath := filepath.Join(storageDir, name) + ".json"
	fmt.Println(filePath)
	return filePath
}

// 处理全部学生信息查询，解析+发送
func doAllstudents(w http.ResponseWriter, r *http.Request) {
	students, err := loadAllStudents()
	if err != nil {
		//http.Error函数的作用的是快速返回错误信息和状态码
		http.Error(w, err.Error(), 500)
		return
	}
	//发送数据给客户端
	err = sendJson(w, students)
	if err != nil {
		fmt.Println("发送信息失败!", err)
		return
	}
	//fmt.Println("1111")
	fmt.Println("发送成功!")

}

// 读取所有学生文件，返回学生列表+错误
func loadAllStudents() ([]*Student, error) {
	// 定义一个切片（类似数组），用来装所有学生对象
	var students []*Student

	// 第一步：遍历存储目录下的所有文件
	// filepath.Walk：递归遍历目录，每个文件/目录都会触发回调函数
	err := filepath.Walk(storageDir, func(path string, info os.FileInfo, err error) error {
		// 如果遍历出错（比如目录权限不足），直接返回
		if err != nil {
			return err
		}

		// 跳过目录，只处理.json文件（避免读取非学生文件）
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".json") {
			return nil
		}

		// 第二步：读取单个JSON文件内容
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("读取文件 %s 失败: %v", path, err)
		}

		// 第三步：JSON转结构体（复用步骤4的反序列化逻辑）
		var stu Student
		if err := json.Unmarshal(data, &stu); err != nil {
			return fmt.Errorf("解析文件 %s 失败: %v", path, err)
		}

		// 第四步：把当前学生添加到列表
		students = append(students, &stu)
		return nil
	})
	// 如果遍历过程出错，返回错误
	if err != nil {
		return nil, err
	}
	// 返回所有学生列表
	return students, nil
}

// 处理单个学生信息查询
func doSingleStudent(w http.ResponseWriter, r *http.Request) {
	//读取单个学生文件，返回单个学生结构体对象+错误
	student, err := LoadStudentFromJSON(getFilePath(" "))
	if err != nil {
		//http.Error函数的作用的是快速返回错误信息和状态码
		http.Error(w, err.Error(), 500)
		return
	}

	//发送数据给客户端
	err = sendJson(w, student)
	if err != nil {
		fmt.Println("发送信息失败!", err)
		return
	}
	fmt.Println("发送成功!")
}

// 读取单个JSON文件并解析为学生结构体对象
func LoadStudentFromJSON(jsonfilename string) (*Student, error) {

	// 读取文件内容
	data, err := os.ReadFile(jsonfilename)
	if err != nil {
		return nil, fmt.Errorf("读取文件内容失败: %v", err)
	}
	// 解析JSON数据到Student结构体,将文本对象转化为内存对象
	var student Student
	err = json.Unmarshal(data, &student)
	if err != nil {
		return nil, fmt.Errorf("解析JSON数据失败: %v", err)
	}
	return &student, nil
}

// 发送json数据给客户端的函数
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
