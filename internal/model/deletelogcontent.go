package model

type ContentInfo struct {
	GameId     int    `form:"game_id"`     // 游戏id
	Tab        string `form:"tab"`         // 玩法tab
	UpdateTime int    `form:"update_time"` // 更新记录的更新时间
}

func DelLogContent(contentInfo ContentInfo) {                                                                            // 操作数据库中的logs表
	logcon := Log{}                                                                                    // 创建一个容器存放数据库中得到的数据
	db.Where("game_id = ? AND tab = ?", contentInfo.GameId, contentInfo.Tab).First(&logcon)            // 查询需要的游戏对应玩法编号（即数据库中的id）
	db.Where("log_id = ? AND update_time = ?", logcon.Id, contentInfo.UpdateTime).Delete(LogContent{}) // 根据log_id update_time删除记录
}
