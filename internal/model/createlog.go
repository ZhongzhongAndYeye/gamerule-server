package model

import (
	"server/utils"

	"github.com/jinzhu/gorm"
)

type Log struct {
	Id      string       `json:"id"`      // 数据库中的id 一个id对应一个游戏的一个tag（前端使用uuid生成传入）
	GameId  int          `json:"game_id"` // 游戏id									（自己规定）
	Name    string                        // 游戏名称
	Tab     string       `json:"tab"`     // 游戏的不同玩法标志
	Version int          `json:"version"` // 版本号
	Logs    []LogContent `json:"logs"`    // 更新 结构体类型 切片
}

type LogContent struct {
	Id         string `json:"id"`          // 数据库中的id							（前端使用uuid生成传入）
	LogId      string `json:"log_id"`      // Log的id
	UpdateTime int64  `json:"update_time"` // 更新时间
	Content    string `json:"content"`     // 更新内容
}

var (
	db *gorm.DB
)

// 拿到在utils中已经初始化好的连接
func init() {
	db = utils.DB
}

// 查询数据库内指定gameid的指定tab是否存在，存在返回false 不存在返回ture
func InquiryGameIdHaveTab(gameid int, tab string) bool {
	logcon := &Log{}
	db.Where("game_id = ? AND tab = ?", gameid, tab).First(logcon)
	if logcon.Id == "" {		// 若是不存在，返回true
		return true
	}else{						// 存在返回false
		return false
	}
}

func CreateLog(log *Log) (createresult bool,msg string){
	db.AutoMigrate(&LogContent{})
	db.AutoMigrate(&Log{})

	result := InquiryGameIdHaveTab(log.GameId,log.Tab)  // 查询需要创建的游戏玩法是否已经存在
	if result == false {
		msg = "此游戏编号下已经存在此玩法，创建失败..."  
		createresult = false
		return
	}else{
		db.Create(log)
		msg = "此游戏编号下的玩法已创建成功..."
		createresult = true
		return
	}
}
