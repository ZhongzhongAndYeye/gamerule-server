package model

func GetLog(gameid int, tab string) (logs []Log, err error, msg string) {
	tx := db // 注意 此处一定要新建会话，避免影响到包内其他使用db的地方
	if gameid != 0 {
		tx = tx.Where("game_id = ?", gameid)
	}
	if tab != "" {
		tx = tx.Where("tab LIKE ?", "%"+tab+"%")
	}
	if err = tx.Find(&logs).Error; err != nil {
		msg = "查找logs表时发生错误..."
	} else {
		msg = "查找logs表成功..."
	}
	return
}
