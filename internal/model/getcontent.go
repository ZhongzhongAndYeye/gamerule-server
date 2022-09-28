package model

func GetContent(id int) (logcontents []LogContent) {
	db.Where("log_id = ?", id).Order("update_time desc").Find(&logcontents) // 根据更新时间降序，更新时间戳越大的数据在越前面
	return
}
