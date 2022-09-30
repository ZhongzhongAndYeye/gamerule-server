package model

func GetTabContent(gameid int, tab string) (data []map[string]interface{}, msg string) {
	now := int64(99) // 获取当前时间，此处为写死的假时间方便调试
	var logs = []Log{}
	var err error
	if logs, err = GetAllLogs(); err != nil {
		msg = "查询数据库时出现错误..."
		return
	}

	for _, log := range logs {
		if log.GameId == gameid && log.Tab == tab {
			for _, content := range log.Contents {
				if content.UpdateTime <= now {
					data = append(data, map[string]interface{}{
						"update_time": content.UpdateTime,
						"content":     content.Content,
					})
				}
			}
		}
	}
	msg = "获取成功..."
	return
}
