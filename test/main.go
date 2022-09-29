package main

import (
	"fmt"
	"server/utils"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// 拿到在utils中已经初始化好的连接
func init() {
	db = utils.DB
}

type Test struct {
	Id int `gorm:"primary_key"`
	Name  string
	Habby  string
}

func main() {
	res := &[]Test{}
	// tx := db
	// tx = tx.Where("name = ?", "b")
    // tx = tx.Where("habby = ?","h")
	// tx.Find(res) 
	fmt.Println("res", res)
	a := 0
	var err error
	if err == nil {
		a = 2
	}
	fmt.Println(a)
}
 