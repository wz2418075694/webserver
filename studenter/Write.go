package studenter

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"webserver/mysql"
)

// 处理创建传入数据
func doWrite(w http.ResponseWriter, r *http.Request) {
	/* io.ReadAll的作用是从实现了io.Reader接口的数据源中(这里是 r)
	读取所有可用的数据，直到遇到结束符（EOF），
	并将读取到的所有数据以 []byte（字节切片）的形式返回。
	*/
	//读取请求体
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("读取请求体失败！" + err.Error()))
		return
	}

	fmt.Println("收到新信息！")

	//将请求体解析成json结构体
	var stu Student //创建学生结构体对象
	err = json.Unmarshal(body, &stu)
	if err != nil {
		http.Error(w, "解析结构体失败!"+err.Error(), 500)
		return
	}

	//结构体解析成功
	//现在得到结构体后，调用刚才的函数，将这个数据写入本地文件中
	err = saveStudentToFile(stu)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	//保存成功，返回给客户端相应
	w.Write([]byte("保存文件成功！"))

	if err := saveStudentToDB(&stu); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("保存DB成功！"))
	//defer result.Close() // 确保语句被关闭
	fmt.Println("db存储成功")

}
func saveStudentToDB(stu *Student) error {
	db := mysql.GetDBHandler() //获取接口
	//执行不返回结果集的SQL语句，并返回执行结果如：影响行数、错误等
	result, err := db.Exec("insert into student(id,name, gender, age, phone, city)"+
		"values(?,?, ?, ?, ?, ?)",
		stu.Id, stu.Name, stu.Gender, stu.Age, stu.Phone, stu.City)
	if err != nil {
		log.Fatal("插入数据失败！", err)
	}
	//获取影响行数
	rowsAffectes, err := result.RowsAffected()
	if err != nil {
		log.Fatal("获取影响行数失败！", err)
	}
	log.Printf("用户插入了%d行数据", rowsAffectes)

	return err
}

// 接收一个学生对象并且把他写入文件中，成功返回nil，失败返回错误信息
func saveStudentToFile(stu Student) error {
	// 第一步：校验必填项（姓名不能为空，否则没法生成文件名）
	if strings.TrimSpace(stu.Name) == "" {
		return fmt.Errorf("学生姓名不能为空")
	}

	// 第二步：唯一文件名（比如：wangzhao.json）
	// 拼接完整路径：存储目录 + 文件名（比如D:\GO\webserver\tmp/wangzhao.json）
	filePath := filepath.Join(storageDir, stu.Name) + ".json"

	// 第三步：把学生结构体转成JSON字符串
	jsonData, err := json.Marshal(stu)
	if err != nil { //序列化失败
		return fmt.Errorf("结构体序列化失败:%v", err)
	}

	// 第四步：把JSON数据写入文件
	// os.WriteFile：创建/覆盖文件，0644可读可写
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}

	// 没有错误，返回nil,表示保存成功
	return nil
}
