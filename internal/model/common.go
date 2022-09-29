package model

import (
	"errors"
	"server/utils"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// 拿到在utils中已经初始化好的连接
func init() {
	db = utils.DB
	db.AutoMigrate(&LogContent{})
	db.AutoMigrate(&Log{})
}

// log基本信息结构体
type Log struct {
	Id            int          `json:"id" gorm:"primary_key"` // 数据库中的id 一个id对应一个游戏的一个tag
	GameId        int          `json:"game_id"`               // 游戏id									（自己规定）
	CustomizeName string       `json:"custmoize_name"`        // 后台自定的名称
	Name          string       `json:"name"`                  // 游戏名称（存在另一张表中）
	Tab           string       `json:"tab"`                   // 游戏的不同玩法标志
	Version       int          `json:"version"`               // 版本号
	Contents      []LogContent `json:"contents"`              // 更新 结构体类型 切片
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
func InquiryGameIdHaveTab(gameid int, tab string) (result bool,err error){
	logcon := &Log{}
	if err = db.Where("game_id = ?", gameid).Where("tab = ?",tab).First(logcon).Error;err != nil{
		if errors.Is(err,gorm.ErrRecordNotFound) {
			err = nil
			result = true
			return 
		}	
		return
	}
	result = false
	return
}

// 查询game_list表内gameid对应的name
func InquiryGameName(gameid int) (name string,err error) {
	nameInfo := &NameInfo{}                                            // 准备game_list表的查询接收容器
	if err=db.Table("game_list").Where("game_id = ?", gameid).First(nameInfo).Error;err!=nil{ // 查询game_id对应的name	
		return
	} 
	name = nameInfo.Name
	return
}
