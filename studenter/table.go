package studenter

import "encoding/json"

type Student struct {
	Id     int    `json:"id"`     //编号
	Name   string `json:"name"`   // 姓名
	Gender string `json:"gender"` // 性别
	Age    int    `json:"age"`    // 年龄
	Phone  string `json:"phone"`  // 电话
	City   string `json:"city"`   // 城市
}

func (s *Student) Encode() ([]byte, error) {
	return json.Marshal(s)
}
