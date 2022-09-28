package model

// log基本信息结构体
type Log struct {
	Id      int          `json:"id" gorm:"primary_key"` // 数据库中的id 一个id对应一个游戏的一个tag
	GameId  int          `json:"game_id"`               // 游戏id									（自己规定）
	Name    string       `json:"name"`                  // 游戏名称
	Tab     string       `json:"tab"`                   // 游戏的不同玩法标志
	Version int          `json:"version"`               // 版本号
	Logs    []LogContent `json:"logs"`                  // 更新 结构体类型 切片
}

// log历史更新记录结构体
type LogContent struct {
	Id         int    `json:"id" gorm:"primary_key"` // 数据库中的id
	LogId      int    `json:"log_id"`                // Log的id
	UpdateTime int64  `json:"update_time"`           // 更新时间
	Content    string `json:"content"`               // 更新内容
}

// 各种game-id对应的name结构体
type NameInfo struct {
	GameId int    `json:"game_id"`
	Name   string `json:"name"`
}

// 查询logs表内指定gameid的指定tab是否存在，存在返回false 不存在返回ture
func InquiryGameIdHaveTab(gameid int, tab string) bool {
	logcon := &Log{}
	db.Where("game_id = ? AND tab = ?", gameid, tab).First(logcon)
	if logcon.Id == 0 { // 若是不存在，返回true
		return true
	} else { // 存在返回false
		return false
	}
}

// 查询game_list表内gameid对应的name
func InquiryGameName(gameid int) (name string) {
	nameInfo := &NameInfo{}                                            // 准备game_list表的查询接收容器
	db.Table("game_list").Where("game_id = ?", gameid).First(nameInfo) // 查询game_id对应的name
	name = nameInfo.Name
	return
}
