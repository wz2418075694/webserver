package mysql

import (
	"database/sql"
	"log"
)

var (
	//创建一个数据库连接池对象
	//操作数据库的唯一入口（执行SQL、事务、配置连接池都靠它）
	db *sql.DB
)

func InitDB() error {
	// 数据库连接信息：用户名:密码@tcp(地址:端口)/数据库名?字符集参数
	dsn := "root:123456@tcp(127.0.0.1:3306)/students?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	//sql.Open()作用是初始化连接池，而非建立连接；
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//完成上面的后，并不会立刻建立网络连接，必须配合Ping()方法验证连接性和可用性
	err = db.Ping()
	if err != nil {
		return err
	}
	// 配置连接池
	db.SetMaxOpenConns(10) // 最大打开连接数
	db.SetMaxIdleConns(5)  // 最大空闲连接数
	log.Println("数据库连接成功")
	return nil
}

func GetDBHandler() *sql.DB {
	return db
}
