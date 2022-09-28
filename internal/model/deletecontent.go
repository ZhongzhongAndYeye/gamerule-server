package model

func DeleteContent(id int) (err error, msg string) {
	err = db.Where("id = ?", id).Delete(&LogContent{}).Error
	if err != nil {
		msg = "删除指定content时发送错误..."
		return
	}
	msg = "已经删除指定的content..."
	return
}
