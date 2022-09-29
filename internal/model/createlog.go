package model

func CreateLog(log *Log) (err error, msg string) {
	var result bool
	if result, err = InquiryGameIdHaveTab(log.GameId, log.Tab); err != nil { // 查询需要创建的游戏玩法是否已经存在
		msg = "查询logs表时出错..."
		return
	}

	if result == false { // 若是已经存在，创建失败
		msg = "此游戏编号下已经存在此玩法，创建失败..."
		return
	} else { // 不存在则创建
		if log.Name, err = InquiryGameName(log.GameId); err != nil {
			msg = "查询game_lists表时出错..."
			return
		}
		if err = db.Create(log).Error; err != nil { // 插入数据进入logs表
			msg = "logs表中创建新数据出错..."
			return
		}
		msg = "此游戏编号下的玩法已创建成功..."
		return
	}
}
