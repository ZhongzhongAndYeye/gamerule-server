package model

import "fmt"

func DeleteLog(id int) (err error, msg string) {
	fmt.Println(id)
	err = db.Where("id = ?", id).Delete(&Log{}).Error // 删除logs表中数据
	if err != nil {
		msg = "删除log过程中发生错误..."
		return
	}
	err = db.Where("log_id = ?", id).Delete(&LogContent{}).Error // 删除log_contents表中指定数据
	if err != nil {
		msg = "删除log_content过程中发生错误..."
		return
	}
	msg = "成功删除指定log以及相应的content..."
	return
}
