package model

import  _ "fmt"

type LogAndContent struct {
	GameId   int    `json:"game_id"`
	Name     string `json:"name"`
	Tab      string `json:"tab"`
	Version  int    `json:"version"`
	Contents []Content
}

type Content struct {
	UpdateTime int64    `json:"update_time"`
	Content    string `json:"content"`
}

func GetLogAndContent(gameid string,tab string) (logAndContent LogAndContent) {
	logcon := Log{}
	db.Where("game_id = ? AND tab = ?",gameid,tab).First(&logcon)
	// fmt.Println(logcon)
	logAndContent.GameId = logcon.GameId
	logAndContent.Name = logcon.Name
	logAndContent.Tab = logcon.Tab
	logAndContent.Version = logcon.Version

	logcontents := []LogContent{}
	db.Order("update_time asc").Where("log_id = ?",logcon.Id).Find(&logcontents)
	// fmt.Println(logcontents) 
	i := len(logcontents)
	logAndContent.Contents = make([]Content,i)
	for index,v := range logcontents {
		logAndContent.Contents[index].UpdateTime = v.UpdateTime
		logAndContent.Contents[index].Content = v.Content
	}
	// fmt.Println(logAndContent.Contents)
	// fmt.Println(logAndContent)
	return
}
