package model

// 获取数据库中两张表全部的数据返回到结构体中
func GetAllLogs() (logs []Log, err error) {

	// 查询数据库中logs表中数据
	if err = db.Find(&logs).Error; err != nil {
		return
	}

	// 将这些数据构造的结构体指针使用map根据其id存储（一定要传入指针！！才能修改原本的切片结构体）
	logIdMap := map[int]*Log{}
	for i := range logs {
		logIdMap[logs[i].Id] = &logs[i]
	}

	// 查询数据库中log_contents表中数据
	logcontents := []LogContent{}
	if err = db.Order("update_time desc").Find(&logcontents).Error; err != nil {
		return
	}

	// 根据标识将每条content分配给对应的log
	for _, content := range logcontents {
		logIdMap[content.LogId].Contents = append(logIdMap[content.LogId].Contents, content)
	}
	return
}

func GetAllTabCondition() (data []map[string]interface{}, msg string) {
	now := int64(55) // 获取当前时间戳，这里写死为50
	var result []Log
	var err error
	if result, err = GetAllLogs(); err != nil {
		msg = "查询数据库时出现错误..."
		return
	}

	for _, log := range result {
		var update_time int64
		for _, content := range log.Contents {
			// 如果更新时间大于当前时间，忽略
			if now < content.UpdateTime {
				continue
			}
			// 如果更新时间小于已经被筛选出来的小于当前时间的更新时间，覆盖
			if update_time < content.UpdateTime {
				update_time = content.UpdateTime
			}
		}
		data = append(data, map[string]interface{}{
			"game_id":     log.GameId,
			"tab":         log.Tab,
			"version":     log.Version,
			"update_time": update_time,
		})
	}
	msg = "获取成功..."
	return
}
