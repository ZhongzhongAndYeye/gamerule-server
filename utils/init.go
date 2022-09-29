package utils

import "gorm.io/gorm"

var (
	DB *gorm.DB
)

// 初始化项目的各种连接
func init(){
	db, _ := Conn("mysql", "root", "123456", "192.168.160.128", "3306", "server")
	DB = db
}