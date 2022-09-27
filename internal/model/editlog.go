package model

func EditLog(log *Log) (changeresult bool, msg string) {

	logcon := &Log{}
	db.Where("id = ?", log.Id).First(logcon)                  // 获取id原来的gameid和tab
	if log.GameId == logcon.GameId && log.Tab == logcon.Tab { // 若是没有编辑原本的游戏id和tab，则直接覆盖
		db.Save(log)
		msg = "此日志已成功编辑..."
		changeresult = true
		return 
	} else {												// 若是改变了原本的游戏id和tab
		result := InquiryGameIdHaveTab(log.GameId, log.Tab) // 查询需要创建的游戏玩法是否已经存在
		if result == false {
			msg = "此游戏编号下已经存在此玩法，编辑失败..."
			changeresult = false
			return
		} else {
			db.Save(log)
			msg = "此日志已成功编辑..."
			changeresult = true
			return
		}
	}
}
