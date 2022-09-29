package utils

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 连接数据库工具
func Conn(dbtype string, username string, password string, url string, port string, d string) (db *gorm.DB, err error) {
	str := username + ":" + password + "@(" + url + ":" + port + ")/" + d + "?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(str),&gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库时出错，错误为:", err)
		return
	}
	fmt.Println("数据库已连接成功...")
	return
}
