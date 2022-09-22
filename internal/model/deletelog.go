package model

type LogInfo struct {
	GameId int    `form:"game_id"`
	Tab    string `form:"tab"`
}

func DelLog(logInfo LogInfo) {
	db.AutoMigrate(&Log{})                                                          // 操作数据库中的logs表
	logcon := Log{}                                                                 // 创建一个容器存放数据库中得到的数据
	db.Where("game_id = ? AND tab = ?", logInfo.GameId, logInfo.Tab).First(&logcon) // 查询需要的游戏对应玩法编号（即数据库中的id）'
	db.Where("log_id = ?", logcon.Id).Delete(LogContent{})							// 将log_contents表中该编号下的所有详情删除
	db.Where("id = ?",logcon.Id).Delete(Log{})										// 将logs表中该编号的游戏玩法删除
}
