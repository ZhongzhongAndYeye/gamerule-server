package model

import _ "fmt"

type NeedLog struct {
	Name    string `json:"name"`
	Tab     string `json:"tab"`
	Version int    `json:"version"`
}

func GetLog() (needlogs []NeedLog) {
	db.AutoMigrate(&Log{}) // 切换到表logs
	logs := []Log{}
	db.Order("game_id asc").Find(&logs) // 根据game_id升序查询记录 放在logs中
	// fmt.Println(logs)
	i := len(logs)						// 获取得到的Log切片长度
	needlogs = make([]NeedLog,i)		// 根据长度开辟NeedLog切片长度
	for index,log := range logs{		// 遍历用logs给needlogs赋值
		needlogs[index].Name = log.Name
		needlogs[index].Tab = log.Tab
		needlogs[index].Version = log.Version
	}
	// fmt.Println(needlogs)
	return								
}
