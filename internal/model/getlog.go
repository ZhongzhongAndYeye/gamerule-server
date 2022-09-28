package model

import "fmt"

func GetLog(gameid int, tab string) (logs []Log) {
	// 四种查询情况
	if gameid == 0 {
		if tab == "" { // 1.传来的game_id为空 tab为空 查询全部
			db.Find(&logs)
			return
		} else { // 2.传来的game_id为空 tab有值 模糊查询tab
			tab_like := fmt.Sprint("%", tab, "%")
			db.Where("tab LIKE ?", tab_like).Find(&logs)
			return
		}
	} else {
		if tab == "" { // 3.传来的game_id有值，tab为空，查询此gameid下的所有tab数据
			db.Where("game_id = ?", gameid).Find(&logs)
			return
		} else { // 4.传来的game_id有值，tab有值，模糊查询此game_id下的tab
			tab_like := fmt.Sprint("%", tab, "%")
			db.Where("game_id = ? AND tab LIKE ?", gameid, tab_like).Find(&logs)
			return
		}
	}
}
