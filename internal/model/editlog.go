package model

func EditLog(log *Log) (msg string, err error) {
	logcon := &Log{}
	if err = db.Where("id = ?", log.Id).First(logcon).Error; err != nil { // 获取id原来的gameid和tab
		msg = "查询logs表时出现错误..."
		return
	}
	if log.GameId == logcon.GameId && log.Tab == logcon.Tab { // 若是没有编辑原本的gameid和tab，则直接覆盖
		if log.Name, err = InquiryGameName(log.GameId); err != nil {
			msg = "查询game_lists表时时出现错误..."
			return
		}
		if err = db.Save(log).Error; err != nil {
			msg = "编辑logs表时出现错误..."
			return
		}
		if err = db.Save(log.Contents).Error;err != nil {
			msg = "编辑log_contents表时出现错误..."
			return
		}
		msg = "此日志已成功编辑..."
		return
	} else { // 若是改变了原本的游戏id和tab
		var result bool
		if result, err = InquiryGameIdHaveTab(log.GameId, log.Tab); err != nil { // 查询需要创建的游戏玩法是否已经存在
			msg = "查询logs表时出现错误..."
			return
		}
		if result == false {
			msg = "此游戏编号下已经存在此玩法，编辑失败..."
			return
		} else {
			if log.Name, err = InquiryGameName(log.GameId); err != nil {
				msg = "查询game_lists表时时出现错误..."
				return
			}
			if err = db.Save(log).Error; err != nil {
				msg = "编辑logs表以及log_contents表时出现错误..."
				return
			}
			if err = db.Save(log.Contents).Error;err != nil {
				msg = "编辑log_contents表时出现错误..."
				return
			}
			msg = "此日志已成功编辑..."
			return
		}
	}
}
