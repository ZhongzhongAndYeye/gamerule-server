package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

type TabInfo struct {
	GameId int    `json:"game_id"`
	Tab    string `json:"tab"`
}

func GetLog(c *gin.Context) {
	tabInfo := &TabInfo{}
	c.ShouldBindJSON(tabInfo)
	logs, err, msg := model.GetLog(tabInfo.GameId, tabInfo.Tab)
	var status int
	if err != nil {
		status = http.StatusBadRequest
	} else {
		status = http.StatusOK
	}
	c.JSON(status, gin.H{
		"msg":  msg,
		"logs": logs,
	})
}
