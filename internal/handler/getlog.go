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
	logs := model.GetLog(tabInfo.GameId, tabInfo.Tab)
	c.JSON(http.StatusOK,logs)
}
