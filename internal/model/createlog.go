package model

import (
	"fmt"
	"server/utils"

	"github.com/jinzhu/gorm"
)

type Log struct {
	Id      int          `json:"id"`                     // 数据库中的id 一个id对应一个游戏的一个tag
	GameId  int          `json:"game_id" form:"game_id"` // 游戏id
	Tab     string       `json:"tab" form:"tab"`         // 游戏的不同玩法标志
	Name    string       `json:"name" form:"name"`       // 游戏名
	Version int          `json:"version" form:"version"` // 版本号
	Count   int          `json:"count" form:"count"`     // 代表当前收到的content切片数量
	Logs    []LogContent // 更新 结构体类型 切片
}

type LogContent struct {
	Id         int    `json:"id"`          // 数据库中的id
	LogId      int    `json:"log_id"`      // ChangeLog的id
	Content    string `json:"content"`     // 更新内容
	UpdateTime int64  `json:"update_time"` // 更新时间
}

var (
	db *gorm.DB
)

//初始化一个数据库连接
func init() {
	d, _ := utils.Conn("mysql", "root", "123456", "192.168.160.128", "3306", "server")
	db = d
}

func CreateLog(log *Log, logcontents []LogContent) {
	db.AutoMigrate(&Log{})                                              // 给数据库根据结构体绑定表
	logcon := Log{}                                                     // 准备一个Log容器，装查询到的数据
	db.Where("game_id=? AND tab=?", log.GameId, log.Tab).First(&logcon) // 查询数据库中是否有对应的gameid以及tab

	if logcon.Id != 0 { // 若是存在 则执行覆盖操作
		logcon.Version = log.Version
		logcon.Count = log.Count
		db.Save(&logcon)
		fmt.Println("覆盖游戏玩法具体信息...")

		id := logcon.Id               // 拿到游戏tab对应的唯一id
		db.AutoMigrate(&LogContent{}) // 操作存放logcontent的表

		var count int
		db.Table("log_contents").Where("log_id=?", id).Count(&count) // 获取游戏的tab下有几条记录
			for _, logcontent := range logcontents { // 遍历收到的切片
				contentcon := LogContent{}                                                           // 准备一个LogContent容器，装查询到的数据
				logcontent.LogId = id                                                                // 将对应的logcontents切片中的logid全部赋值上id值以形成联系
				db.Where("log_id=? AND update_time=?", id, logcontent.UpdateTime).First(&contentcon) // 查询游戏具体玩法的某个时间是否已经有更新详情
				if contentcon.Id != 0 {                                                              // 若是有 则覆盖
					contentcon.Content = logcontent.Content
					db.Save(&contentcon)
					fmt.Println("编号为", id, "的tab，更新时间戳为", logcontent.UpdateTime, "的玩法更新已经覆盖原玩法")
				} else { // 若是没有 则新建
					db.Create(&logcontent)
					fmt.Println("编号为", id, "的tab，更新时间戳为", logcontent.UpdateTime, "的玩法更新已经新建")
				}
			}
	} else { // 不存在则直接插入
		db.Create(&log)
		id := logcon.Id
		fmt.Println("新建游戏玩法具体信息...")
		db.AutoMigrate(&LogContent{})
		for _, logcontent := range logcontents {
			logcontent.LogId = id
			db.Create(&logcontent)
			fmt.Println("编号为", id, "的tab，更新时间戳为", logcontent.UpdateTime, "的玩法更新已经新建")
		}
	}
}
