package model

import (
	"server/utils"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

// 拿到在utils中已经初始化好的连接
func init() {
	db = utils.DB
}

func CreateLog(log *Log) (createresult bool, msg string) {
	db.AutoMigrate(&LogContent{})
	db.AutoMigrate(&Log{})

	result := InquiryGameIdHaveTab(log.GameId, log.Tab) // 查询需要创建的游戏玩法是否已经存在
	if result == false {                                // 若是已经存在，创建失败
		msg = "此游戏编号下已经存在此玩法，创建失败..."
		createresult = false
		return
	} else { // 不存在则创建
		log.Name = InquiryGameName(log.GameId)
		db.Create(log) // 插入数据进入logs表
		msg = "此游戏编号下的玩法已创建成功..."
		createresult = true
		return
	}
}
