package model

func GetContent(id int) (logcontents []LogContent,msg string,err error) {
	if err = db.Where("log_id = ?", id).Order("update_time desc").Find(&logcontents).Error;err!=nil{// 根据更新时间降序，更新时间戳越大的数据在越前面
		msg = "查询log_contents表时发生错误..."
		return
	} 
	msg = "查询log_contents表成功..."
	return
}
